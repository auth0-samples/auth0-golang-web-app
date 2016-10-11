package main

import (
        "github.com/auth0-samples/auth0-golang-web-app/01-Login/app"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Init()
	StartServer()

}
