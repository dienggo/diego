package controllers

import "github.com/gin-gonic/gin"

type Pong struct{}

func (controller Pong) Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong 2",
	})
}
