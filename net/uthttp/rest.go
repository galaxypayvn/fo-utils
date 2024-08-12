package uthttp

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/valyala/fasthttp"
	"gitlab.com/goxp/cloud0/logger"
)

const (
	ContentTypeJSON           = "application/json"
	ContentTypeXML            = "application/xml"
	ContentTypeFormURLEncoded = "application/x-www-form-urlencoded"
)

type APIRequest struct {
	Method      string            // HTTP method
	URL         string            // URL to send the request to
	Header      map[string]string // HTTP headers
	Param       map[string]string // Query parameters
	Body        interface{}       // Body of the request
	LogTag      string            // Tag to use for logging
	ContentType string            // Content-Type of the request body
}

func SendRestAPI[SuccessResponse any, ErrorResponse any](ctx context.Context, req APIRequest) (*SuccessResponse, *ErrorResponse, error) {
	var (
		log = logger.WithCtx(ctx, req.LogTag)
	)

	httpRequest := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpRequest)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	httpRequest.SetRequestURI(req.URL)
	httpRequest.Header.SetMethod(req.Method)

	for key, value := range req.Header {
		httpRequest.Header.Set(key, value)
	}

	queryArgs := httpRequest.URI().QueryArgs()
	for key, value := range req.Param {
		queryArgs.Set(key, value)
	}

	var (
		bodyData []byte
		err      error
	)
	if req.Body != nil {
		// Set the Content-Type header if not already set
		if req.ContentType != "" {
			httpRequest.Header.SetContentType(req.ContentType)
		} else {
			// Default to "application/json" if Content-Type is not specified
			httpRequest.Header.SetContentType(ContentTypeJSON)
		}

		switch req.ContentType {
		case ContentTypeJSON:
			bodyData, err = json.Marshal(req.Body)
		case ContentTypeXML:
			bodyData, err = xml.Marshal(req.Body)
		case ContentTypeFormURLEncoded:
			bodyData, err = structToURLValues(req.Body)
		default:
			bodyData, err = json.Marshal(req.Body)
		}
		if err != nil {
			log.WithError(err).Error("error marshalling httpRequest body")
			return nil, nil, err
		}

		httpRequest.SetBody(bodyData)
	}
	log.Infof("api: %v header: %v requestData: %v", req.URL, req.Header, string(bodyData))

	if err := fasthttp.Do(httpRequest, resp); err != nil {
		log.WithError(err).Error("error sending httpRequest")
		return nil, nil, err
	}

	statusCode := resp.StatusCode()
	// Handle error response
	if statusCode < 200 || statusCode >= 300 {
		var errorResponse ErrorResponse
		if err := json.Unmarshal(resp.Body(), &errorResponse); err != nil {
			log.WithError(err).Error("error unmarshalling error response body")
			return nil, nil, fmt.Errorf("%w: %w", ErrUnmarshalResponse, err)
		}
		log.Infof("api: %v statusCode: %v responseData: %+v", req.URL, statusCode, errorResponse)
		return nil, &errorResponse, nil
	}

	// Handle successful response
	var result SuccessResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		log.WithError(err).Error("error unmarshalling success response body")
		return nil, nil, fmt.Errorf("%w: %w", ErrUnmarshalResponse, err)
	}
	log.Infof("api: %v statusCode: %v responseData: %+v", req.URL, statusCode, result)

	return &result, nil, nil
}
