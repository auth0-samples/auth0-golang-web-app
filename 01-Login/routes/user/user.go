package user

import (
	"net/http"

	"github.com/gorilla/sessions"

	"01-Login/routes/templates"
)

func Handler(store *sessions.FilesystemStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		templates.RenderTemplate(w, "user", session.Values["profile"])
	})
}
