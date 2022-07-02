package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type web struct{}

func (web) Do(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		fmt.Fprintf(c.Writer, "Hello world!")
	})
}

func Init() web {
	return web{}
}