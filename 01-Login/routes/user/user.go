package user

import (
        "github.com/auth0/auth0-golang/examples/regular-web-app/app"
        templates "github.com/auth0/auth0-golang/examples/regular-web-app/routes"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := app.GlobalSessions.SessionStart(w, r)
	defer session.SessionRelease(w)

	templates.RenderTemplate(w, "user", session.Get("profile"))
}
