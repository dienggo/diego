package main

import (
	"github.com/joho/godotenv"
	"go_base_project/routes"
	"os"
	"time"
)

func main() {
	println("App started")
	time.LoadLocation(os.Getenv("TIME_ZONE"))
	godotenv.Load()
	routes.Init()
}