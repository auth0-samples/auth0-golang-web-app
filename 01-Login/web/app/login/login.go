package login

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"01-Login/platform/authenticator"
	"01-Login/web/app/pkce"
)

// Handler for our login.
func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		verifier, err := pkce.RandomVerifier(32)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Save the state inside the session.
		session := sessions.Default(ctx)
		session.Set("state", state)
		session.Set("verifier", verifier)
		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		authCodeOptions := pkce.Sha256Challenge(verifier)
		ctx.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(state, authCodeOptions...))
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
