package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hrabit64/shortlink/app/schema"
	"github.com/hrabit64/shortlink/app/service"
	"github.com/hrabit64/shortlink/app/userErrors"
)

func UpdateUser(c *gin.Context) {
	var req schema.UserUpdateRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
			"error":   err.Error(),
		})
		return
	}

	// 사용자 정보 업데이트 로직
	res, err := service.UpdateUser(req)

	if err != nil {
		var businessErr *userErrors.BusinessError
		if errors.As(err, &businessErr) {
			c.JSON(
				businessErr.Status,
				gin.H{
					"status":  businessErr.Status,
					"message": businessErr.Message,
				},
			)
			return
		}
		panic(err)
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "사용자 정보가 업데이트 되었습니다.",
		"result":  res,
	})
}
