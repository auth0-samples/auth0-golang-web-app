package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"net/http"
	"net/url"
	"os"

	"github.com/auth0/go-auth0/v2/authentication/oauth"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
	gob.Register(map[string]interface{}{})
}

func initSessionStore() {
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
		Secure:   false, // Set to true in production (requires HTTPS)
		SameSite: http.SameSiteLaxMode,
	}
}

// HomeHandler renders the home page or redirects to /user if already logged in.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth-session")
	if session.Values["profile"] != nil {
		http.Redirect(w, r, "/user", http.StatusSeeOther)
		return
	}
	if err := templates.ExecuteTemplate(w, "home.html", nil); err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
}

// LoginHandler redirects the user to Auth0's Universal Login page.
func LoginHandler(auth *Authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state, err := generateRandomState()
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		session, _ := store.Get(r, "auth-session")
		session.Values["state"] = state
		if err := session.Save(r, w); err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, auth.AuthorizationURL(state), http.StatusTemporaryRedirect)
	}
}

// CallbackHandler handles the callback from Auth0 after authentication.
func CallbackHandler(auth *Authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "auth-session")

		// Verify the state parameter to prevent CSRF attacks.
		if r.URL.Query().Get("state") != session.Values["state"] {
			http.Error(w, "Invalid state parameter", http.StatusBadRequest)
			return
		}

		// Exchange the authorization code for tokens.
		tokenSet, err := auth.OAuth.LoginWithAuthCode(r.Context(), oauth.LoginWithAuthCodeRequest{
			Code:        r.URL.Query().Get("code"),
			RedirectURI: auth.CallbackURL,
		}, oauth.IDTokenValidationOptions{})
		if err != nil {
			http.Error(w, "Failed to exchange authorization code for token", http.StatusUnauthorized)
			return
		}

		// Retrieve the user's profile information.
		userInfo, err := auth.UserInfo(r.Context(), tokenSet.AccessToken)
		if err != nil {
			http.Error(w, "Failed to get user info", http.StatusInternalServerError)
			return
		}

		session.Values["access_token"] = tokenSet.AccessToken
		session.Values["profile"] = map[string]interface{}{
			"nickname": userInfo.Nickname,
			"name":     userInfo.Name,
			"picture":  userInfo.Picture,
			"email":    userInfo.Email,
		}
		if err := session.Save(r, w); err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/user", http.StatusTemporaryRedirect)
	}
}

// UserHandler displays the authenticated user's profile.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth-session")
	profile, ok := session.Values["profile"].(map[string]interface{})
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := templates.ExecuteTemplate(w, "user.html", profile); err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
}

// LogoutHandler clears the session and redirects to Auth0's logout endpoint.
func LogoutHandler(auth *Authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "auth-session")
		session.Options.MaxAge = -1
		session.Save(r, w)

		logoutURL, _ := url.Parse("https://" + auth.Domain + "/v2/logout")
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		returnTo, _ := url.Parse(scheme + "://" + r.Host)
		params := url.Values{}
		params.Add("returnTo", returnTo.String())
		params.Add("client_id", auth.ClientID)
		logoutURL.RawQuery = params.Encode()

		http.Redirect(w, r, logoutURL.String(), http.StatusTemporaryRedirect)
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
