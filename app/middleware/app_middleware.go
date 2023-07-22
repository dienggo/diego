package middleware

import (
	"github.com/gin-gonic/gin"
	"go_base_project/config"
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
