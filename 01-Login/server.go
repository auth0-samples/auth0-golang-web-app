package main

import (
        "github.com/auth0-samples/auth0-golang-web-app/01-Login/routes/callback"
        "github.com/auth0-samples/auth0-golang-web-app/01-Login/routes/home"
        "github.com/auth0-samples/auth0-golang-web-app/01-Login/routes/middlewares"
        "github.com/auth0-samples/auth0-golang-web-app/01-Login/routes/user"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
