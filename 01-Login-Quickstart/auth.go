package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/auth0/go-auth0/v2/authentication"
)

// Authenticator wraps the go-auth0 authentication client.
type Authenticator struct {
	*authentication.Authentication
	Domain      string
	ClientID    string
	CallbackURL string
}

// NewAuthenticator creates and configures a new Authenticator.
func NewAuthenticator() (*Authenticator, error) {
	domain := os.Getenv("AUTH0_DOMAIN")
	clientID := os.Getenv("AUTH0_CLIENT_ID")
	clientSecret := os.Getenv("AUTH0_CLIENT_SECRET")
	callbackURL := os.Getenv("AUTH0_CALLBACK_URL")

	authClient, err := authentication.New(
		context.Background(),
		domain,
		authentication.WithClientID(clientID),
		authentication.WithClientSecret(clientSecret),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize authentication client: %w", err)
	}

	return &Authenticator{
		Authentication: authClient,
		Domain:         domain,
		ClientID:       clientID,
		CallbackURL:    callbackURL,
	}, nil
}

// AuthorizationURL builds the /authorize URL to redirect users
// to Auth0's Universal Login page.
func (a *Authenticator) AuthorizationURL(state string) string {
	u, _ := url.Parse("https://" + a.Domain + "/authorize")
	params := url.Values{
		"response_type": {"code"},
		"client_id":     {a.ClientID},
		"redirect_uri":  {a.CallbackURL},
		"scope":         {"openid profile email"},
		"state":         {state},
	}
	u.RawQuery = params.Encode()
	return u.String()
}
