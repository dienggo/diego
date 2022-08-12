package base

import (
	"github.com/gin-gonic/gin"
	"net/url"
)

type Controller struct {}

func (controller Controller) Request(c *gin.Context) request {
	return request{c: c}
}

// Request is Data type of Request data
type Request struct {
	Param  map[string]any `json:"param"`
	Header map[string]any `json:"header"`
}

// GetParam is function to get value of request parameter by key
func (req Request) GetParam(key string) any {
	return req.Param[key]
}

// GetHeader is function to get value of request header by key
func (req Request) GetHeader(key string) any {
	return req.Header[key]
}

// Private request to wrapping the method
type request struct {
	c *gin.Context
}

// Data is simplified function to get data from gin context
func (req request) Data() Request {
	data := Request{}
	param := map[string]any{}
	header := map[string]any{}

	for k, v := range req.c.Request.Header { header[k] = v[0] }

	// Handling GET method
	u, _ := url.Parse(req.c.Request.URL.String())
	queryParams := u.Query()
	for k, v := range queryParams { param[k] = v[0] }

	// Handling POST method
	if req.c.Request.Method == "POST" { req.c.MultipartForm() }
	for k, v := range req.c.Request.PostForm { param[k] = v[0] }
	var myMap map[string]any
	req.c.ShouldBindJSON(&myMap)
	for k, v := range myMap { param[k] = v }

	// Bind data on model
	data.Param = param
	data.Header = header
	return data
}

// Value is function to get value by key
func(req request) Value(key string) any {
	return req.Data().GetParam(key)
}