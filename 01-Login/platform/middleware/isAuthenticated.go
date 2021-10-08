package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func IsAuthenticated(
	store *sessions.FilesystemStore,
) func(
	w http.ResponseWriter,
	r *http.Request,
	next http.HandlerFunc,
) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		session, err := store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, ok := session.Values["profile"]; !ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			next(w, r)
		}
	}
}
