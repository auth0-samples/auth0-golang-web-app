package app

import (
	"encoding/gob"
	"log"

	"github.com/joho/godotenv"
	"github.com/gorilla/sessions"
)

var (
	Store *sessions.FilesystemStore
)

func Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Print(err.Error())
		return err
	}

	Store = sessions.NewFilesystemStore("", []byte("something-very-secret"))
	gob.Register(map[string]interface{}{})
	return nil
}
