package main

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"

	"01-Login/platform/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load the env vars: %v", err)
	}

	store := sessions.NewFilesystemStore("", []byte("something-very-secret"))
	gob.Register(map[string]interface{}{})

	rtr := router.New(store)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("there was an error with the http server: %v", err)
	}
}
