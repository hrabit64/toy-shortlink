package handlers

import "github.com/gin-gonic/gin"

func GetLoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
	return
}

func GetMainPage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
	return
}
