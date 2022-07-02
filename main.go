package main

import (
	"github.com/joho/godotenv"
	"go_base_project/routes"
)

func main() {
	godotenv.Load()
	routes.Init()
}