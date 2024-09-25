package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ErrorHandlingMiddleware(c *gin.Context) {
	defer func() {
		// 패닉 복구
		if err := recover(); err != nil {

			log.Println(err)

			// 패닉이 발생하면 500 에러와 함께 응답
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			c.Abort()
		}
	}()

	// 다음 핸들러로 넘어감
	c.Next()

	// c.Errors에 저장된 에러가 있을 경우
	if len(c.Errors) > 0 {
		// 첫 번째 에러만 가져옴
		err := c.Errors[0]
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
