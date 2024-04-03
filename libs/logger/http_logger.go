package logger

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"bitbucket.org/finesys/finesys-utility/libs/serror"
)

type (
	Request struct {
		Url     string      `json:"url"`
		Method  string      `json:"method"`
		Headers interface{} `json:"header"`
		Body    interface{} `json:"body"`
	}
	Response struct {
		Code    int         `json:"status"`
		Headers interface{} `json:"header"`
		Body    interface{} `json:"body"`
	}
	HttpLogDataObject struct {
		Type     string   `json:"type"`
		Request  Request  `json:"request"`
		Response Response `json:"response"`
	}
)

func CreateLogData(respHttp *http.Response) (resp HttpLogDataObject, errx serror.SError) {
	var (
		reqBody, resBody interface{}
		body             []byte
		err              error
	)

	if respHttp.Request.Body != nil {
		body, err = ioutil.ReadAll(respHttp.Request.Body)
		if err != nil {
			errx = serror.NewFromErrorc(err, "[internal][CreateLogParams] while reading request body")
			return
		}

		err = json.Unmarshal(body, &reqBody)
		if err != nil {
			reqBody = string(body)
		}
	}

	if respHttp.Body != nil {
		body, err = ioutil.ReadAll(respHttp.Body)
		if err != nil {
			errx = serror.NewFromErrorc(err, "[internal][createLogParams] while reading response body")
			return
		}
		err = json.Unmarshal(body, &resBody)
		if err != nil {
			resBody = string(body)
		}
	}

	defer respHttp.Body.Close()

	resp = HttpLogDataObject{
		Type: "http",
		Request: Request{
			Url:     respHttp.Request.Host + respHttp.Request.URL.RequestURI(),
			Method:  respHttp.Request.Method,
			Headers: respHttp.Request.Header,
			Body:    reqBody,
		},
		Response: Response{
			Code:    respHttp.StatusCode,
			Headers: respHttp.Header,
			Body:    resBody,
		},
	}

	return
}
