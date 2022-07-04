package main

import (
	"github.com/joho/godotenv"
	"go_base_project/app/base"
	"go_base_project/routes"
)

func main() {
	godotenv.Load()
	dbInstant := base.DbInstantiation()
	defer dbInstant.Close()
	routes.Init()
}