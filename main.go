package main

import (
	"github.com/joho/godotenv"
	"go_base_project/routes"
)

func main() {
	println("App started")
	godotenv.Load()
	routes.Init()
}