package uthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"gitlab.com/goxp/cloud0/logger"
)

var (
	ErrUnmarshalResponse = errors.New("failed to unmarshal response")
)

// HTTPRequest represents an HTTP request.
type HTTPRequest struct {
	Method string            // HTTP method
	URL    string            // URL to send the request to
	Header map[string]string // HTTP headers
	Body   interface{}       // Body of the request
	LogTag string            // Tag to use for logging
}

type HTTPResponse[T any] struct {
	StatusCode int
	Header     map[string][]string
	Body       T
}

type Config struct {
	Timeout time.Duration
}

type Options struct {
	DisallowUnknownFields bool
}

func NewHTTPClient(cfg Config) *http.Client {
	client := &http.Client{Timeout: cfg.Timeout}
	return client
}

func SendHTTPRequest[T any](ctx context.Context, client *http.Client, httpReq HTTPRequest, opts ...Options) (HTTPResponse[T], error) {
	var (
		log       = logger.WithCtx(ctx, httpReq.LogTag)
		bodyBytes []byte
		resp      *http.Response
		err       error
	)

	var res HTTPResponse[T]
	// Marshal request body
	if httpReq.Body != nil {
		bodyBytes, err = json.Marshal(httpReq.Body)
		if err != nil {
			log.WithError(err).Error("error marshaling request body")
			return res, err
		}
		log.Infof("api: %v header: %v responseData: %v", httpReq.URL, httpReq.Header, string(bodyBytes))
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, httpReq.Method, httpReq.URL, bytes.NewReader(bodyBytes))
	if err != nil {
		log.WithError(err).Error("error creating HTTP request")
		return res, err
	}

	// Add headers
	req.Header.Set("x-request-id", fmt.Sprint(ctx.Value("x-request-id")))
	req.Header.Set("Content-Type", "application/json")
	for k, v := range httpReq.Header {
		req.Header.Set(k, v)
	}

	// Send HTTP request
	resp, err = client.Do(req)
	if err != nil {
		log.WithError(err).Error("error sending HTTP request")
		return res, err
	}
	defer resp.Body.Close()
	res.StatusCode = resp.StatusCode
	res.Header = resp.Header

	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != http.StatusAccepted {
		responseByte, _ := io.ReadAll(resp.Body)
		log.Infof("api: %v statusCode: %v responseData: %v", httpReq.URL, resp.StatusCode, string(responseByte))
	}

	dec := json.NewDecoder(resp.Body)
	if len(opts) > 0 {
		opt := opts[0]
		if opt.DisallowUnknownFields {
			dec.DisallowUnknownFields()
		}
	}

	var body T
	// Unmarshal response
	err = dec.Decode(&body)
	if err != nil {
		log.WithError(err).Errorf("api: %v error unmarshalling response", httpReq.URL)
		return res, fmt.Errorf("%w: %w", ErrUnmarshalResponse, err)
	}
	log.Infof("api: %v responseData: %+v", httpReq.URL, body)
	res.Body = body

	return res, nil
}

func MakeURL(baseURL, path string) string {
	baseURL = strings.TrimSuffix(baseURL, "/")
	path = strings.TrimPrefix(path, "/")

	return fmt.Sprintf("%s/%s", baseURL, path)
}
