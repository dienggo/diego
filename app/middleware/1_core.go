package middleware

import "github.com/gin-gonic/gin"

type IMiddleware interface {
	Handle(c *gin.Context)
}

func Add(middlewares ...IMiddleware) (a add) {
	a = add{}
	var ms []IMiddleware
	for _, middleware := range middlewares {
		ms = append(ms, middleware)
	}
	a.middlewares = ms
	return a
}

type add struct {
	middlewares []IMiddleware
}

// Handle : handle middleware registered
func (receiver add) Handle(c *gin.Context) {
	for _, middleware := range receiver.middlewares {
		middleware.Handle(c)
	}
}
