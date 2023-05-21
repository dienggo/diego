package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Web struct{}

func (w Web) Do(route *gin.Engine) {

	// example test web
	route.GET("/", func(c *gin.Context) {
		fmt.Fprintf(c.Writer, "Hello world!")
	})
}
