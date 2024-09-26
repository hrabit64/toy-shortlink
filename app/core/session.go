package core

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"os"
)

var store *sessions.CookieStore

// var store = sessions.NewCookieStore([]byte("123"), []byte())
func GetSessionStore() *sessions.CookieStore {
	return store
}

func GetSession(c *gin.Context) (*sessions.Session, error) {
	return store.Get(c.Request, "session")
}

func GetIsAuthenticated(c *gin.Context) bool {
	session, _ := GetSession(c)
	authenticated := session.Values["authenticated"]
	return authenticated != nil && authenticated == true
}

func SetIsAuthenticated(c *gin.Context, value bool) {
	session, _ := GetSession(c)
	session.Values["authenticated"] = value
	session.Save(c.Request, c.Writer)
}

func InitSessionStore() {
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")), []byte(os.Getenv("SESSION_ENC_KEY")))

	store.Options = &sessions.Options{
		MaxAge:   2 * 60 * 60,
		HttpOnly: true,
		Path:     "/",
	}
}
