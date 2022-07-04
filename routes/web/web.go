package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_base_project/app/base"
)

type web struct{}

func (web) Do(router *base.Router) {
	router.GET("/", func(c *gin.Context) {
		fmt.Fprintf(c.Writer, "Hello world!")
	})
}

func Init() web {
	return web{}
}