package messagecode

import (
	"context"
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
}

type Client struct {
	redisCli *redis.Client
	cfg      Config
}

type strapiMessageCodeResp struct {
	Data []strapiMessageCode `json:"data"`
	Meta strapiMeta          `json:"meta"`
}

type strapiMessageCode struct {
	ID      int    `json:"id"`
	Code    int    `json:"code"`
	Locale  string `json:"locale"`
	Message string `json:"message"`
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

func NewClient(cfg Config) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPwd,
		DB:       cfg.RedisDB,
	})

	return &Client{
		redisCli: rdb,
		cfg:      cfg,
	}
}

func (c *Client) GetMessage(ctx context.Context, messageGroup int, messageCodes []int) (map[string]string, error) {
	keys := MakeKeys(supportLocales, messageCodes)
	res := c.redisCli.MGet(ctx, keys...)
	err := res.Err()
	if err != nil {
		return nil, err
	}

	messages := res.Val()
	messageMap := make(map[string]string, len(keys))
	for idx, key := range keys {
		if messages[idx] != nil {
			var ok bool
			messageMap[key], ok = messages[idx].(string)
			if !ok {
				return nil, errors.New("message is not a string")
			}
		} else {
			messageMap, err = c.getMessagesFromStrapi(ctx, messageGroup)
			if err != nil {
				return nil, err
			}

			cmdRes := c.redisCli.MSet(ctx, messageMap)
			err = cmdRes.Err()
			if err != nil {
				return nil, err
			}

			return messageMap, nil
		}

	}

	return messageMap, nil
}

func (c *Client) getMessagesFromStrapi(ctx context.Context, messageGroup int) (map[string]string, error) {
	res := map[string]string{}
	messageCodes, err := c.getStrapiMessageCodes(ctx, messageGroup)
	if err != nil {
		return nil, err
	}
	for _, messCode := range messageCodes {
		res[MakeKey(messCode.Locale, messCode.Code)] = messCode.Message
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
		queryVals.Set("pagination[pageSize]", "1")
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
		opt := uthttp.HTTPOptions{
			Timeout: 3 * time.Second,
		}

		resp, err := uthttp.SendHTTPRequest[strapiMessageCodeResp](ctx, req, opt)
		if err != nil {
			return nil, err
		}

		messageCodes = append(messageCodes, resp.Data...)

		totalPage = resp.Meta.Pagination.PageCount
		page++
	}

	return messageCodes, nil
}

func MakeKeys(locales []string, messageCodes []int) []string {
	keys := make([]string, 0, len(locales)*len(messageCodes))
	for _, locale := range locales {
		for _, code := range messageCodes {
			keys = append(keys, MakeKey(locale, code))
		}
	}

	return keys
}

func MakeKey(locale string, messageCode int) string {
	return fmt.Sprintf("message_code:%s:%d", locale, messageCode)
}
