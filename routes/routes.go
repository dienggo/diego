package routes

import (
	"github.com/gin-gonic/gin"
	"go_base_project/routes/api"
	"go_base_project/routes/web"
)

type RouteInterface interface {
	Do(router *gin.Engine)
}

func Init() {
	router := gin.New()
	api.Init().Do(router)
	web.Init().Do(router)
	router.Run()
}