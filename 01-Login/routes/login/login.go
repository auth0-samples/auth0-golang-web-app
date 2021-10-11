package login

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gorilla/sessions"

	"01-Login/auth"
)

func Handler(store *sessions.FilesystemStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate random state
		b := make([]byte, 32)
		_, err := rand.Read(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		state := base64.StdEncoding.EncodeToString(b)

		session, err := store.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["state"] = state
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		authenticator, err := auth.NewAuthenticator()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)

	})
}
