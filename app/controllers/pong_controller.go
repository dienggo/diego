package controllers

import "github.com/gin-gonic/gin"

type PongController struct{}

func (controller PongController) Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong 2",
	})
}
