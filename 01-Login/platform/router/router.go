package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/urfave/negroni"

	"01-Login/platform/middleware"
	"01-Login/routes/callback"
	"01-Login/routes/home"
	"01-Login/routes/login"
	"01-Login/routes/logout"
	"01-Login/routes/user"
)

func New(store *sessions.FilesystemStore) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", home.HomeHandler)

	router.Handle("/login", login.Handler(store))
	router.Handle("/logout", logout.Handler())
	router.Handle("/callback", callback.Handler(store))

	router.Handle(
		"/user",
		negroni.New(
			negroni.HandlerFunc(middleware.IsAuthenticated(store)),
			negroni.Wrap(user.Handler(store)),
		),
	)

	router.PathPrefix("/public/").Handler(
		http.StripPrefix(
			"/public/",
			http.FileServer(http.Dir("public/")),
		),
	)

	return router
}
