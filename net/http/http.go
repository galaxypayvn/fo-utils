package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"gitlab.com/goxp/cloud0/logger"
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

// SendHTTPRequest sends an HTTP request and returns the response.
// It takes a context, an HTTPRequest, and HTTPOptions.
// It returns the response as an interface{} and any error encountered.
func SendHTTPRequest[T any](ctx context.Context, httpReq HTTPRequest, options HTTPOptions) (T, error) {
	var (
		log       = logger.WithCtx(ctx, options.LogTag)
		bodyBytes []byte
		respBytes []byte
		resp      *http.Response
		err       error
	)

	var res T
	// Marshal request body
	if httpReq.Body != nil {
		bodyBytes, err = json.Marshal(httpReq.Body)
		if err != nil {
			if options.ErrorLog {
				log.WithError(err).Error("error marshaling request body")
			}
			return res, err
		}
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, httpReq.Method, httpReq.URL, bytes.NewReader(bodyBytes))
	if err != nil {
		if options.ErrorLog {
			log.WithError(err).Error("error creating HTTP request")
		}
		return res, err
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	for k, v := range httpReq.Header {
		req.Header.Set(k, v)
	}

	// Send HTTP request
	client := &http.Client{Timeout: options.Timeout}
	resp, err = client.Do(req)
	if err != nil {
		if options.ErrorLog {
			log.WithError(err).Error("error sending HTTP request")
		}
		return res, err
	}
	defer resp.Body.Close()

	// Read response body
	respBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		if options.ErrorLog {
			log.WithError(err).Error("error reading response body")
		}
		return res, err
	}
	if options.ResponseLog {
		log.Infof("api: %v responseData: %v", httpReq.URL, string(respBytes))
	}

	// Unmarshal response
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		if options.ErrorLog {
			log.WithError(err).Error("error unmarshalling response")
		}
		return res, err
	}

	return res, nil
}
