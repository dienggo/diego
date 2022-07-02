package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"testing"
)

func TestDatabase(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
		return
	}
	db := Database()
	println("database name : ",db.Name)
}