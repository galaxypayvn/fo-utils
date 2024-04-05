package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"gitlab.com/goxp/cloud0/logger"
	"io"
	"net"
	"net/http"
	"time"
)

// HTTPRequest represents an HTTP request.
type HTTPRequest struct {
	Method string            // HTTP method
	URL    string            // URL to send the request to
	Header map[string]string // HTTP headers
	Body   interface{}       // Body of the request
}

// HTTPOptions represents options for sending an HTTP request.
type HTTPOptions struct {
	Timeout     time.Duration // Timeout for the request
	ErrorLog    bool          // Whether to log errors
	ResponseLog bool          // Whether to log responses
	LogTag      string        // Tag to use for logging
}

func setHTTPOptionsDefaults(options *HTTPOptions) *HTTPOptions {
	if options == nil {
		return &HTTPOptions{
			Timeout:     30 * time.Second,
			ErrorLog:    true,
			ResponseLog: true,
			LogTag:      "SendHTTPRequest",
		}
	}

	if options.Timeout == 0 {
		options.Timeout = 30 * time.Second
	}

	if options.LogTag == "" {
		options.LogTag = "SendHTTPRequest"
	}

	return options
}

// SendHTTPRequest sends an HTTP request and returns the response.
// It takes a context, an HTTPRequest, and HTTPOptions.
// It returns the response as an interface{} and any error encountered.
func SendHTTPRequest(ctx context.Context, httpReq HTTPRequest, options *HTTPOptions) (interface{}, error) {
	var (
		log       = logger.WithCtx(ctx, options.LogTag)
		bodyBytes []byte
		respBytes []byte
		resp      *http.Response
		err       error
	)
	options = setHTTPOptionsDefaults(options)

	// Marshal request body
	if httpReq.Body == nil {
		bodyBytes = nil
	} else {
		bodyBytes, err = json.Marshal(httpReq.Body)
		if err != nil && options.ErrorLog {
			log.Infof("HTTP error: %v errorMsg: %v", "marshal body error", err.Error())
		}
	}
	if options.ResponseLog {
		log.Infof("Request data: method: %v api: %v header: %v body: %v", httpReq.Method, httpReq.URL, httpReq.Header, string(bodyBytes))
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, httpReq.Method, httpReq.URL, bytes.NewReader(bodyBytes))
	if err != nil {
		if options.ErrorLog {
			log.Infof("HTTP error: %v errorMsg: %v", "http.NewRequestWithContext error", err.Error())
		}
		return nil, err
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(HeaderRequestID, ctx.Value(HeaderRequestID).(string))
	for k, v := range httpReq.Header {
		req.Header.Set(k, v)
	}

	req.Close = true

	// Send HTTP request
	client := &http.Client{Timeout: options.Timeout}
	resp, err = client.Do(req)
	if err != nil {
		if options.ErrorLog {
			log.Infof("HTTP error: %v errorMsg: %v", "client.Do error", err.Error())
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != http.StatusAccepted {
		var e net.Error
		if errors.As(err, &e) && e.Timeout() {
			log.Infof("HTTP error: %v error: %v", "time_out error", e.Error())
			return nil, e
		}

		responseByte, _ := io.ReadAll(resp.Body)
		log.Infof("HTTP error: %v statusCode: %v error: %v responseData: %v", "internal error", resp.StatusCode, err.Error(), string(respBytes))
		return nil, errors.New(string(responseByte))
	}

	// Read response body
	respBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		if options.ErrorLog {
			log.Infof("HTTP error: %v statusCode: %v error: %v response: %v", "read response error", resp.StatusCode, err.Error(), string(respBytes))
		}
		return nil, err
	}
	if options.ResponseLog {
		log.Infof("HTTP response: method: %v api: %v body: %v", httpReq.Method, httpReq.URL, string(respBytes))
	}

	// Unmarshal response
	var response interface{}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		if options.ErrorLog {
			log.Infof("HTTP error: %v errorMsg: %v", "unmarshal response error", err.Error())
		}
		return nil, err
	}

	return response, nil
}
