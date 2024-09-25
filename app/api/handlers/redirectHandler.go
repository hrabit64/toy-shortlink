package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hrabit64/shortlink/app/service"
	"github.com/hrabit64/shortlink/app/userErrors"
)

func ToOriginalUrl(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	res, err := service.LookUpItemByShortPath(shortUrl)

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

	c.Redirect(302, res.OriginURL)
	return
}
