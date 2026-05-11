package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var templates *template.Template

func main() {
	// .env file is optional; in Docker, env vars come from --env-file flag.
	godotenv.Load()

	initSessionStore()

	auth, err := NewAuthenticator()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	templates = template.Must(template.ParseGlob("templates/*.html"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/login", LoginHandler(auth))
	mux.HandleFunc("/callback", CallbackHandler(auth))
	mux.HandleFunc("/user", UserHandler)
	mux.HandleFunc("/logout", LogoutHandler(auth))

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", mux); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
