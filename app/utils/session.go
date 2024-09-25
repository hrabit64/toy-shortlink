package utils

import (
	"github.com/gorilla/sessions"
	"os"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

func GetSessionStore() *sessions.CookieStore {
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 1, // 1시간
		HttpOnly: true,
	}
	return store
}
