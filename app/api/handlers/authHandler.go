package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hrabit64/shortlink/app/service"
	"github.com/hrabit64/shortlink/app/utils"
)

func ProcessLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := service.GetUserById(username)

	if err != nil {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "사용자 ID가 올바르지 않습니다.",
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
	store := utils.GetSessionStore()
	session, _ := store.Get(c.Request, "session")
	session.Values["username"] = username
	session.Save(c.Request, c.Writer)

	c.JSON(200, gin.H{
		"status":  200,
		"message": "로그인 성공",
	})

	return

}

func ProcessLogout(c *gin.Context) {
	store := utils.GetSessionStore()
	session, _ := store.Get(c.Request, "session")
	delete(session.Values, "username")
	session.Save(c.Request, c.Writer)

	c.JSON(200, gin.H{
		"status":  200,
		"message": "로그아웃 성공",
	})

	return
}
