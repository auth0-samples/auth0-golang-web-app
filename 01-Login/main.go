package main

import (
        "github.com/auth0/auth0-golang/examples/regular-web-app/app"
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
