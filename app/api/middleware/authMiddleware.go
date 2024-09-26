package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hrabit64/shortlink/app/core"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	store := core.GetSessionStore()

	session, _ := store.Get(c.Request, "session")
	user := session.Values["user"]

	if user == nil {
		c.Redirect(http.StatusFound, "/login?error=로그인이 필요합니다!")
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()
}
