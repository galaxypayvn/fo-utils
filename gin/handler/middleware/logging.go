package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"gitlab.com/goxp/cloud0/logger"
)

func LoggingRequest(isProduction bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.WithCtx(c, "Request Detail")
		defer func() {
			if r := recover(); r != nil {
				log.Error(r)
				debug.PrintStack()
				panic(r)
			}
		}()
		r := c.Request
		//params := c.Request.URL.Query()
		header := c.Request.Header
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Errorf("Error reading request body: %v", err.Error())
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		var obj interface{}
		_ = json.Unmarshal(buf, &obj)

		data, _ := json.Marshal(obj)
		if !isProduction {
			//log.WithField("body", fmt.Sprintf("%s", data)).WithField("header", header).Info("uri: ", c.Request.RequestURI)
			log.Info("uri: ", c.Request.RequestURI, " header: ", header, " body: ", fmt.Sprintf("%s", data))
		} else {
			//log.WithField("body", fmt.Sprintf("%s", data)).Info("uri: ", c.Request.RequestURI)
			appVersion := c.Request.Header.Get("x-current-version")
			log.Info("uri: ", c.Request.RequestURI, " body: ", string(data), " x-current-version: ", appVersion)
		}
		reader := io.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = reader
		c.Next()
	}
}
