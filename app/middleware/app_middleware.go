package middleware

import (
	"github.com/dienggo/diego/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type App struct{}

func (App) Handle(c *gin.Context) {
	appKey := c.GetHeader("app-key")
	if config.App().Key != appKey {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "you can't do this!",
		})
		return
	}
}
