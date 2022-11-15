package main

import (
	"go_base_project/app/base"
	"go_base_project/routes"
	"log"

	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	log.Println("App started")

	// load env configuration
	godotenv.Load()

	// set development mode
	gin.SetMode(gin.TestMode)

	// set database
	db := base.OpenDB()
	if db == nil {
		log.Println("Failed to connect to database")
	} else {
		log.Println("Connected to database")
	}

	routes.Init()
}
