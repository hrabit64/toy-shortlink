package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hrabit64/shortlink/app/core"
	"log"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	session, _ := core.GetSession(c)
	log.Println("values : ", session.Values)
	user := session.Values["username"]

	if user == nil {
		c.Redirect(http.StatusFound, "/login?error=로그인이 필요합니다!")
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()
}
