package base

import "github.com/gin-gonic/gin"

type Controller struct {}

func (controller Controller) Request(c *gin.Context) request {
	return request{c: c}
}

type request struct {
	c *gin.Context
}

// Value is function to get value by key
func(req request) Value(key string) string {
	switch req.c.Request.Method {
	case "GET":
		return req.c.Query(key)
	case "POST":
		dataPost := req.c.PostForm(key)
		if dataPost == "" {
			var myMap = make(map[string]string)
			req.c.BindJSON(&myMap)
			return myMap[key]
		}
		return dataPost
	default:
		return ""
	}
}