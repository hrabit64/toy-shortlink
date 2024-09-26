package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hrabit64/shortlink/app/core"
	"github.com/hrabit64/shortlink/app/service"
	"github.com/hrabit64/shortlink/app/utils"
	"log"
	"net/http"
)

func ProcessLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := service.GetUserById(username)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(401, gin.H{
				"status":  401,
				"message": "사용자 ID가 올바르지 않습니다.",
			})
			return
		}

		c.JSON(500, gin.H{
			"status":  500,
			"message": "서버에 문제가 발생했습니다.",
		})

		return
	}

	if !utils.CheckPasswordHash(password, user.Pw) {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "비밀번호가 올바르지 않습니다.",
		})
		return
	}
	session, _ := core.GetSession(c)
	session.Values["authenticated"] = true
	err = session.Save(c.Request, c.Writer)

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"status":  500,
			"message": "세션 저장에 실패했습니다.",
		})
		return
	}

	c.Header("HX-Redirect", "/")
	c.JSON(200, gin.H{
		"message": "Login successful",
	})
	return

}

func ProcessLogout(c *gin.Context) {
	core.SetIsAuthenticated(c, false)
	c.Redirect(http.StatusFound, "/login?error=로그인이 필요합니다.")

	return
}
