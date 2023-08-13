package controllers

import "github.com/gin-gonic/gin"

type Pong struct{}

func (ctrl Pong) Main(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong!",
	})
}
