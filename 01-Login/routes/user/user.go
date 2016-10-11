package user

import (
        "github.com/auth0-samples/auth0-golang-web-app/01-Login/app"
        templates "github.com/auth0-samples/auth0-golang-web-app/01-Login/routes"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := app.GlobalSessions.SessionStart(w, r)
	defer session.SessionRelease(w)

	templates.RenderTemplate(w, "user", session.Get("profile"))
}
