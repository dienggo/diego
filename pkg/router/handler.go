package router

import (
	"bytes"
	"encoding/json"
	"github.com/dienggo/diego/pkg/logger"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func handler(context *gin.Context) {
	timeStart := time.Now()
	body, _ := io.ReadAll(context.Request.Body)
	context.Request.Body = io.NopCloser(bytes.NewReader(body))
	context.Next()

	request, _ := json.Marshal(context.Request.URL)
	header, _ := json.Marshal(context.Request.Header)

	// Calculate response time
	timeEnd := time.Now()
	responseTime := timeEnd.Sub(timeStart)

	wrapped := map[string]any{
		"status":            context.Writer.Status(),
		"url":               context.Request.URL,
		"request":           string(request),
		"body":              string(body),
		"header":            string(header),
		"client_ip":         context.ClientIP(),
		"remote_ip":         context.RemoteIP(),
		"response_time_mcs": responseTime.Microseconds(),
	}

	var loggerFields []logger.Field

	for key, val := range wrapped {
		loggerFields = append(loggerFields, logger.SetField(key, val))
	}

	if context.Writer.Status() >= 400 {
		logger.Error("Error HTTP Observer "+context.Request.URL.Path, loggerFields...)
	} else {
		logger.Info("Info HTTP Observer "+context.Request.URL.Path, loggerFields...)
	}
}
