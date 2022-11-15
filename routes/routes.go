package routes

import (
	"go_base_project/app/base"
	"go_base_project/routes/api"
	"go_base_project/routes/web"
	"os"
)

type RouteInterface interface {
	Do(router *base.Router)
}

func Init() {
	router := base.NewRouter()
	api.Init().Do(router)
	web.Init().Do(router)

	router.Run(":" + os.Getenv("APP_PORT"))
}
