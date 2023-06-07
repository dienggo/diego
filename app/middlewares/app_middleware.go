package middlewares

import (
	"github.com/gin-gonic/gin"
	"go_base_project/config"
	"net/http"
)

// AppMiddleware to restrict api connection
func AppMiddleware(c *gin.Context) {
	appKey := c.GetHeader("app-key")
	if config.App().Key != appKey {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "you can't do this!",
		})
		return
	}
}
