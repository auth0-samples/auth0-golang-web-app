package app

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
)

var (
	Store *sessions.CookieStore
)

func Init() error {
	Store = sessions.NewCookieStore([]byte("something-very-secret"))
	gob.Register(map[string]interface{}{})
	return nil
}
