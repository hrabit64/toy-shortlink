package core

import (
	"github.com/gorilla/sessions"
	"os"
)

var store *sessions.CookieStore

// var store = sessions.NewCookieStore([]byte("123"), []byte())
func GetSessionStore() *sessions.CookieStore {
	return store
}

func InitSessionStore() {
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")), []byte(os.Getenv("SESSION_ENC_KEY")))

	store.Options = &sessions.Options{
		MaxAge:   2 * 60 * 60,
		HttpOnly: true,
		Path:     "/",
	}
}
