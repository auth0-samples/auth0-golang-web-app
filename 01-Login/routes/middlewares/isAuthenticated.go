package middlewares

import (
        "github.com/auth0/auth0-golang/examples/regular-web-app/app"
	"net/http"
)

func IsAuthenticated(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	session, _ := app.GlobalSessions.SessionStart(w, r)
	defer session.SessionRelease(w)
	if session.Get("profile") == nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		next(w, r)
	}
}
