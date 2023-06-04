package pkce

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/oauth2"
)

const (
	codeChallengeKey       = "code_challenge"
	codeChallengeMethodKey = "code_challenge_method"
	codeVerifierKey        = "code_verifier"

	challengeMethodSha256 = "S256"
	challengeMethodPlain  = "plain"
)

func RandomVerifier(length int) (string, error) {
	// RFC7636 4.1: Result must be between 43 and 128 characters.
	if length < 32 || length > 96 {
		return "", fmt.Errorf("expected length between 32 and 96, got %d", length)
	}

	buf := make([]byte, length)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}

func Sha256Challenge(verifier string) []oauth2.AuthCodeOption {
	sha := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(sha[:])

	return challengeOptions(challengeMethodSha256, challenge)
}

func PlainChallenge(verifier string) []oauth2.AuthCodeOption {
	return challengeOptions(challengeMethodPlain, verifier)
}

func challengeOptions(method string, challenge string) []oauth2.AuthCodeOption {
	return []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam(codeChallengeMethodKey, method),
		oauth2.SetAuthURLParam(codeChallengeKey, challenge),
	}
}

func ExchangeVerifier(verifier string) []oauth2.AuthCodeOption {
	return []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam(codeVerifierKey, verifier),
	}
}
