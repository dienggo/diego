package base

import "github.com/gin-gonic/gin"

type handlerFunc func(ctx *gin.Context)

func methodHandler(ctx *gin.Context, handle handlerFunc) {
	defer DbInstantiation().Close()
	handle(ctx)
}

type RouterInterface interface {
	GET(string, handlerFunc) gin.IRoutes
	POST(string, handlerFunc) gin.IRoutes
	DELETE(string, handlerFunc) gin.IRoutes
	PATCH(string, handlerFunc) gin.IRoutes
	PUT(string, handlerFunc) gin.IRoutes
}

func NewRouter() *Router {
	r := new(Router)
	r.gin = gin.New()
	return r
}

type Router struct{ gin *gin.Engine }

func (r Router) Group(relativePath string, handlers ...gin.HandlerFunc) *Router {
	group := r.gin.Group(relativePath, handlers...)
	newRoute := gin.Default()
	newRoute.RouterGroup = *group
	rr := new(Router)
	rr.gin = newRoute
	return rr
}

func (r Router) Run(addr ...string) (err error) {
	return r.gin.Run(addr...)
}

func (r Router) GET(s string, h handlerFunc) gin.IRoutes {
	return r.gin.GET(s, func(context *gin.Context) {methodHandler(context, h)})
}

func (r Router) POST(s string, h handlerFunc) gin.IRoutes {
	return r.gin.POST(s, func(context *gin.Context) {methodHandler(context, h)})
}

func (r Router) DELETE(s string, h handlerFunc) gin.IRoutes {
	return r.gin.DELETE(s, func(context *gin.Context) {methodHandler(context, h)})
}

func (r Router) PATCH(s string, h handlerFunc) gin.IRoutes {
	return r.gin.PATCH(s, func(context *gin.Context) {methodHandler(context, h)})
}

func (r Router) PUT(s string, h handlerFunc) gin.IRoutes {
	return r.gin.PUT(s, func(context *gin.Context) {methodHandler(context, h)})
}
