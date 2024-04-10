package messagecode

import (
	"context"
	"encoding/json"
	"errors"
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
)

var supportLocales = []string{"en", "vi"}

type Config struct {
	RedisAddr            string
	RedisPwd             string
	RedisDB              int
	StrapiMessageCodeURL string
	StrapiToken          string
	MessageGroup         int
	MessageCodes         []int
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
	err := client.loadMessageCode(ctx, cfg.MessageGroup, cfg.MessageCodes)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) GetMessage(locale string, code int) string {
	return c.messageMap[makeKey(locale, code)].Message
}

func (c *Client) GetHTTPCode(locale string, code int) int {
	httpCode := c.messageMap[makeKey(locale, code)].HTTPCode
	if httpCode == 0 {
		return fallbackMessageCodeToHTTPCode(code)
	}

	return httpCode
}

func (c *Client) loadMessageCode(ctx context.Context, messageGroup int, messageCodes []int) error {
	keys := makeKeys(supportLocales, messageCodes)
	res := c.redisCli.MGet(ctx, keys...)
	err := res.Err()
	if err != nil {
		return err
	}

	messages := res.Val()
	for idx, key := range keys {
		if messages[idx] != nil {
			var ok bool
			rawMessageCode, ok := messages[idx].(string)
			if !ok {
				return errors.New("message is not a string")
			}

			var mc messageCode
			err := json.Unmarshal([]byte(rawMessageCode), &mc)
			if err != nil {
				return err
			}

			c.messageMap[key] = mc
		} else {
			messageMap, err := c.getMessageMapFromStrapi(ctx, messageGroup)
			if err != nil {
				return err
			}

			c.messageMap = messageMap

			byteMap, err := messageMapToByteMap(messageMap)
			if err != nil {
				return err
			}

			cmdRes := c.redisCli.MSet(ctx, byteMap)
			err = cmdRes.Err()
			if err != nil {
				return err
			}

			return nil
		}

	}

	return nil
}

func (c *Client) getMessageMapFromStrapi(ctx context.Context, messageGroup int) (map[string]messageCode, error) {
	res := map[string]messageCode{}
	messageCodes, err := c.getStrapiMessageCodes(ctx, messageGroup)
	if err != nil {
		return nil, err
	}
	for _, messCode := range messageCodes {
		res[makeKey(messCode.Locale, messCode.Code)] = messageCode{
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
		queryVals.Set("filters[$or][0][code][$startsWithi]", strconv.Itoa(messageGroup))
		queryVals.Set("filters[$or][1][code][$startsWithi]", strconv.Itoa(generalGroup))
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

func messageMapToByteMap(messageMap map[string]messageCode) (map[string]any, error) {
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

func makeKeys(locales []string, messageCodes []int) []string {
	keys := make([]string, 0, len(locales)*len(messageCodes))
	for _, locale := range locales {
		for _, code := range messageCodes {
			keys = append(keys, makeKey(locale, code))
		}
	}

	return keys
}

func makeKey(locale string, messageCode int) string {
	return fmt.Sprintf("messagecode:%s:%d", locale, messageCode)
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
