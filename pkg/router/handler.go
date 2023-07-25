package router

import (
	"encoding/json"
	"github.com/dienggo/diego/pkg/logger"
	"github.com/gin-gonic/gin"
	"io"
)

func handler(context *gin.Context) {
	if context.Writer.Status() >= 400 {
		body, _ := io.ReadAll(context.Request.Body)
		request, _ := json.Marshal(&context.Request.URL)
		header, _ := json.Marshal(&context.Request.Header)

		wrapped := map[string]any{
			"request":   string(request),
			"body":      string(body),
			"header":    string(header),
			"client_ip": context.ClientIP(),
			"remote_ip": context.RemoteIP(),
		}

		logger.Error("Error HTTP Observer "+context.Request.URL.Path,
			logger.SetField("wrapped", wrapped),
			logger.SetField("url", context.Request.URL),
			logger.SetField("status", context.Writer.Status()),
		)
	}
}
