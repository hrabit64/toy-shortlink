package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hrabit64/shortlink/app/utils"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	store := utils.GetSessionStore()

	session, _ := store.Get(c.Request, "session")
	user := session.Values["user"]

	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()
}
