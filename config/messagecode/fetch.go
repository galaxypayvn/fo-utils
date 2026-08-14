package messagecode

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code.finan.one/finan-one-be/fo-utils/net/uthttp"
	redis "github.com/redis/go-redis/v9"
)

const generalGroup = 10

type messageCode struct {
	HTTPCode int    `json:"http_code"`
	Message  string `json:"messasge"`
}

type fileMessageCode struct {
	ID       int    `json:"id"`
	Code     int    `json:"code"`
	Locale   string `json:"locale"`
	Message  string `json:"message"`
	HTTPCode int    `json:"http_code"`
}

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

	if err := client.LoadMessageCode(context.Background(), cfg.MessageGroup...); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) GetMessage(locale string, code int) string {
	messCode, ok := c.messageMap[makeFieldKey(locale, code)]
	if !ok {
		return getDefaultLocaleMessage(locale)
	}

	return messCode.Message
}

func (c *Client) GetHTTPCode(locale string, code int) int {
	messCode, ok := c.messageMap[makeFieldKey(locale, code)]
	if !ok {
		return fallbackMessageCodeToHTTPCode(code)
	}

	return messCode.HTTPCode
}

// LoadMessageCode loads catalog entries from the embedded JSON file.
// Duplicate locale+code pairs keep the last occurrence, matching previous Strapi merge order.
func (c *Client) LoadMessageCode(_ context.Context, messageGroups ...int) error {
	codes, err := loadEmbeddedMessageCodes()
	if err != nil {
		return err
	}

	groups := append([]int{}, messageGroups...)
	groups = append(groups, generalGroup)

	messageCodeMap := make(map[string]messageCode, len(codes))
	for _, messCode := range codes {
		if !codeBelongsToGroups(messCode.Code, groups) {
			continue
		}
		messageCodeMap[makeFieldKey(messCode.Locale, messCode.Code)] = messageCode{
			HTTPCode: messCode.HTTPCode,
			Message:  messCode.Message,
		}
	}

	c.mergeMessageCodesMap(messageCodeMap)
	return nil
}

func loadEmbeddedMessageCodes() ([]fileMessageCode, error) {
	var codes []fileMessageCode
	if err := json.Unmarshal(messageCodesJSON, &codes); err != nil {
		return nil, fmt.Errorf("parse embedded message_codes.json: %w", err)
	}
	return codes, nil
}

func codeBelongsToGroups(code int, groups []int) bool {
	codeStr := strconv.Itoa(code)
	for _, group := range groups {
		if strings.HasPrefix(codeStr, strconv.Itoa(group)) {
			return true
		}
	}
	return false
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

func (c *Client) mergeMessageCodesMap(messageMap map[string]messageCode) {
	for key, val := range messageMap {
		c.messageMap[key] = val
	}
}

func (c *Client) PublishMessageCode(ctx context.Context, req CreateMessageCodeReq) (interface{}, error) {

	var unifiedResponse interface{}
	if req.EnMessage == "" || req.ViMessage == "" {
		return nil, errors.New("missing English message or Vietnamese message")
	}

	// Prepare the initial POST request
	req.Locale = "en"
	req.Message = req.EnMessage
	httpReq := uthttp.HTTPRequest{
		Method: http.MethodPost,
		URL:    c.cfg.StrapiMessageCodeURL,
		Header: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", c.cfg.StrapiToken),
		},
		Body:   req,
		LogTag: "PublishMessageCode",
	}

	client := uthttp.NewHTTPClient(uthttp.Config{
		Timeout: 3 * time.Second,
	})

	// Send the initial POST request
	res, err := uthttp.SendHTTPRequest[json.RawMessage](ctx, client, httpReq, uthttp.DefaultOptions())
	if err != nil {
		return unifiedResponse, err
	}

	if res.StatusCode >= 400 {
		var errorResponse ErrorResponse
		if err := json.Unmarshal(res.Body, &errorResponse); err != nil {
			return unifiedResponse, err
		}
		unifiedResponse = &errorResponse
	} else {
		var successResponse SuccessResponse
		if err := json.Unmarshal(res.Body, &successResponse); err != nil {
			return unifiedResponse, err
		}
		unifiedResponse = &successResponse

		// Prepare the PUT request to update the English message
		req.Locale = "vi"
		req.Message = req.ViMessage
		req.Localizations = []int{successResponse.Data.ID}
		httpReq = uthttp.HTTPRequest{
			Method: "POST",
			URL:    c.cfg.StrapiMessageCodeURL,
			Header: map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", c.cfg.StrapiToken),
			},
			Body:   req,
			LogTag: "PublishMessageCode",
		}

		res, err = uthttp.SendHTTPRequest[json.RawMessage](ctx, client, httpReq, uthttp.DefaultOptions())
		if err != nil {
			return unifiedResponse, err
		}

		if res.StatusCode >= 400 {
			var errorResponse ErrorResponse
			if err := json.Unmarshal(res.Body, &errorResponse); err != nil {
				return unifiedResponse, err
			}
			unifiedResponse = &errorResponse
		} else {
			var updateResponse SuccessResponse
			if err := json.Unmarshal(res.Body, &updateResponse); err != nil {
				return unifiedResponse, err
			}
			unifiedResponse = &updateResponse
		}
	}

	return unifiedResponse, nil
}
