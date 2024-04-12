package messagecode

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"code.finan.cc/finan-one-be/fo-utils/net/uthttp"
	redis "github.com/redis/go-redis/v9"
)

const (
	generalGroup = 10

	messageGroupEmptyKey = "empty"
)

type Config struct {
	RedisAddr            string
	RedisPwd             string
	RedisDB              int
	StrapiMessageCodeURL string
	StrapiToken          string
	MessageGroup         []int
}

type Client struct {
	redisCli   *redis.Client
	cfg        Config
	messageMap map[string]messageCode
}

type messageCode struct {
	HTTPCode int    `json:"http_code"`
	Message  string `json:"messasge"`
}

type strapiMessageCodeResp struct {
	Data []strapiMessageCode `json:"data"`
	Meta strapiMeta          `json:"meta"`
}

type strapiMessageCode struct {
	ID       int    `json:"id"`
	Code     int    `json:"code"`
	Locale   string `json:"locale"`
	Message  string `json:"message"`
	HTTPCode int    `json:"http_code"`
}

type strapiMeta struct {
	Pagination strapiPagination `json:"pagination"`
}

type strapiPagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	PageCount int `json:"pageCount"`
	Total     int `json:"total"`
}

func NewClient(cfg Config) (*Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPwd,
		DB:       cfg.RedisDB,
	})

	client := &Client{
		redisCli:   rdb,
		cfg:        cfg,
		messageMap: map[string]messageCode{},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.loadMessageCode(ctx, cfg.MessageGroup...)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) GetMessage(locale string, code int) string {
	return c.messageMap[makeFieldKey(locale, code)].Message
}

func (c *Client) GetHTTPCode(locale string, code int) int {
	httpCode := c.messageMap[makeFieldKey(locale, code)].HTTPCode
	if httpCode == 0 {
		return fallbackMessageCodeToHTTPCode(code)
	}

	return httpCode
}

func (c *Client) loadMessageCode(ctx context.Context, messageGroups ...int) error {
	messageGroups = append(messageGroups, generalGroup)
	for _, group := range messageGroups {
		key := makeHashKey(group)
		messageGroupRes := c.redisCli.HGetAll(ctx, key)
		err := messageGroupRes.Err()
		if err != nil {
			return err
		}

		messageStrMap := messageGroupRes.Val()
		if len(messageStrMap) != 0 {
			_, ok := messageStrMap[messageGroupEmptyKey]
			if ok {
				continue
			}
			messageCodeMap, err := byteMapToMessageCodeMap(messageStrMap)
			if err != nil {
				return err
			}

			c.mergeMessageCodesMap(messageCodeMap)
		} else {
			messageCodeMap, err := c.getMessageGroupMapFromStrapi(ctx, group)
			if err != nil {
				return err
			}

			if len(messageCodeMap) == 0 {
				messageCodeMap[messageGroupEmptyKey] = messageCode{}
			}

			anyMap, err := messageMapToAnyMap(messageCodeMap)
			if err != nil {
				return err
			}

			cmdRes := c.redisCli.HMSet(ctx, key, anyMap)
			err = cmdRes.Err()
			if err != nil {
				return err
			}

			c.mergeMessageCodesMap(messageCodeMap)
		}
	}

	return nil
}

func (c *Client) getMessageGroupMapFromStrapi(ctx context.Context, messageGroup int) (map[string]messageCode, error) {
	res := map[string]messageCode{}
	messageCodes, err := c.getStrapiMessageCodes(ctx, messageGroup)
	if err != nil {
		return nil, err
	}
	for _, messCode := range messageCodes {
		res[makeFieldKey(messCode.Locale, messCode.Code)] = messageCode{
			HTTPCode: messCode.HTTPCode,
			Message:  messCode.Message,
		}
	}

	return res, nil
}

func (c *Client) getStrapiMessageCodes(ctx context.Context, messageGroup int) ([]strapiMessageCode, error) {
	totalPage := 1
	page := 1
	uri, err := url.Parse(c.cfg.StrapiMessageCodeURL)
	if err != nil {
		return nil, err
	}
	var messageCodes []strapiMessageCode
	for page <= totalPage {
		queryVals := uri.Query()
		queryVals.Set("pagination[page]", fmt.Sprintf("%d", page))
		queryVals.Set("pagination[pageSize]", "300")
		queryVals.Set("locale", "all")
		queryVals.Set("filters[code][$startsWithi]", strconv.Itoa(messageGroup))
		uri.RawQuery = queryVals.Encode()

		req := uthttp.HTTPRequest{
			Method: http.MethodGet,
			URL:    uri.String(),
			Header: map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", c.cfg.StrapiToken),
			},
		}
		cfg := uthttp.Config{
			Timeout: 3 * time.Second,
		}

		client := uthttp.NewHTTPClient(cfg)

		resp, err := uthttp.SendHTTPRequest[strapiMessageCodeResp](ctx, client, req)
		if err != nil {
			return nil, err
		}

		messageCodes = append(messageCodes, resp.Data...)

		totalPage = resp.Meta.Pagination.PageCount
		page++
	}

	return messageCodes, nil
}

func (c *Client) mergeMessageCodesMap(messageMap map[string]messageCode) {
	for key, val := range messageMap {
		c.messageMap[key] = val
	}
}

func messageMapToAnyMap(messageMap map[string]messageCode) (map[string]any, error) {
	byteMap := make(map[string]any, len(messageMap))

	for key, val := range messageMap {
		blob, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}

		byteMap[key] = blob
	}

	return byteMap, nil
}

func byteMapToMessageCodeMap(byteMap map[string]string) (map[string]messageCode, error) {
	messsageCodeMap := make(map[string]messageCode, len(byteMap))
	for key, val := range byteMap {
		var messCode messageCode
		err := json.Unmarshal([]byte(val), &messCode)
		if err != nil {
			return nil, err
		}

		messsageCodeMap[key] = messCode
	}

	return messsageCodeMap, nil
}

func makeHashKey(messageGroup int) string {
	return fmt.Sprintf("messagegroup:%d", messageGroup)
}

func makeFieldKey(locale string, messageCode int) string {
	return fmt.Sprintf("%s:%d", locale, messageCode)
}

func fallbackMessageCodeToHTTPCode(code int) int {
	messCodeStr := fmt.Sprintf("%d", code)

	if len(messCodeStr) != 6 {
		return http.StatusInternalServerError
	}

	switch messCodeStr[2] {
	case '2':
		return http.StatusOK
	case '4':
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
