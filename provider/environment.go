package provider

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func loadEnv(wd string, increment int) {
	split := strings.Split(wd, "/")
	newSplit := split[:len(split)-increment]
	joined := strings.Join(newSplit, "/")
	err := godotenv.Load(joined + "/.env")
	if err != nil && increment > 10 {
		panic("can't load the env : " + err.Error())
	} else if err != nil && increment <= 10 {
		increment++
		loadEnv(wd, increment)
	} else if err == nil {
		println("Successfully loaded .env with total tries :", increment+1)
	} else {
		panic("can't load the env with no specific reason")
	}
}

func LoadEnv() {
	wd, _ := os.Getwd()
	loadEnv(wd, 0)
}
