package handlers

import "github.com/gin-gonic/gin"

func GetLoginPage(c *gin.Context) {
	c.HTML(200, "login.tmpl", gin.H{})
	return
}
