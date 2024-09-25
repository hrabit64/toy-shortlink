package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hrabit64/shortlink/app/schema"
	"github.com/hrabit64/shortlink/app/service"
	"github.com/hrabit64/shortlink/app/userErrors"
	"github.com/hrabit64/shortlink/app/utils"
	"strconv"
)

func GetItems(c *gin.Context) {

	pageable, err := utils.NewPageable(c)

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

	items, err := service.GetAllItems(pageable)

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

	c.JSON(200, items)
}

func GetItem(c *gin.Context) {

	reqId := c.Param("id")

	id, err := strconv.Atoi(reqId)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
		})
		return
	}

	item, err := service.GetItemById(id)

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

	c.JSON(200, schema.NewItemResponse(item))
}

func CreatePermItem(c *gin.Context) {

	var item schema.PermItemCreateRequest

	err := c.ShouldBindJSON(&item)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
			"error":   err.Error(),
		})
		return
	}

	permItem := schema.PermItemCreateData{
		OriginURL: item.OriginURL,
		ShortPath: item.ShortPath,
	}

	result, err := service.CreatePermItem(permItem)

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
		"message": "성공적으로 등록되었습니다.",
		"id":      result,
	})
}

func CreateTempItem(c *gin.Context) {

	var item schema.TempItemCreateRequest

	err := c.ShouldBindJSON(&item)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
			"error":   err.Error(),
		})
		return
	}

	tempItem := schema.TempItemCreateData{
		OriginURL: item.OriginURL,
		ShortPath: item.ShortPath,
		ExpSec:    item.ExpSec,
	}

	result, err := service.CreateTempItem(tempItem)

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
		"message": "성공적으로 등록되었습니다.",
		"id":      result,
	})
}

func CreateCountItem(c *gin.Context) {

	var item schema.CountItemCreateRequest

	err := c.ShouldBindJSON(&item)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
			"error":   err.Error(),
		})
		return
	}

	countItem := schema.CountItemCreateData{
		OriginURL: item.OriginURL,
		ShortPath: item.ShortPath,
		InitCount: item.InitCount,
	}

	result, err := service.CreateCountItem(countItem)

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
		"message": "성공적으로 등록되었습니다.",
		"id":      result,
	})
}

func UpdatePermItem(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
		})
		return
	}

	var item schema.PermItemUpdateRequest

	err = c.ShouldBindJSON(&item)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
			"error":   err.Error(),
		})
		return
	}

	permItem := schema.PermItemUpdateData{
		Id:        id,
		OriginURL: item.OriginURL,
		ShortPath: item.ShortPath,
	}

	err = service.UpdatePermItem(permItem)

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
		"message": "성공적으로 수정되었습니다.",
	})
}

func UpdateTempItem(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
		})
		return
	}

	var item schema.TempItemUpdateRequest

	err = c.ShouldBindJSON(&item)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
			"error":   err.Error(),
		})
		return
	}

	tempItem := schema.TempItemUpdateData{
		Id:        id,
		OriginURL: item.OriginURL,
		ShortPath: item.ShortPath,
		ExpSec:    item.ExpSec,
	}

	err = service.UpdateTempItem(tempItem)

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
		"message": "성공적으로 수정되었습니다.",
	})
}

func UpdateCountItem(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
		})
		return
	}

	var item schema.CountItemUpdateRequest

	err = c.ShouldBindJSON(&item)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
			"error":   err.Error(),
		})
		return
	}

	countItem := schema.CountItemUpdateData{
		Id:        id,
		OriginURL: item.OriginURL,
		ShortPath: item.ShortPath,
		InitCount: item.InitCount,
	}

	err = service.UpdateCountItem(countItem)

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
		"message": "성공적으로 수정되었습니다.",
	})
}

func ConvertItem(c *gin.Context) {

	targetType := c.Query("type")

	if targetType != "perm" && targetType != "temp" && targetType != "count" {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
		})
		return
	}

	if targetType == "perm" {
		convertToPermItem(c, err, id)
		return
	}

	if targetType == "temp" {
		convertToTempItem(c, err, id)
		return
	}

	if targetType == "count" {
		convertToCountItem(c, err, id)
		return
	}
}

func convertToCountItem(c *gin.Context, err error, id int) {
	var item schema.CountItemCreateRequest

	err = c.ShouldBindJSON(&item)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
			"error":   err.Error(),
		})
		return
	}

	countItem := &schema.CountItemConvertedData{
		Id:        id,
		OriginURL: item.OriginURL,
		ShortPath: item.ShortPath,
		InitCount: item.InitCount,
	}

	res, err := service.ConvertToCountItem(countItem)

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
		"message": "성공적으로 변환되었습니다.",
		"result":  schema.NewItemResponse(res),
	})
	return
}

func convertToTempItem(c *gin.Context, err error, id int) {
	var item schema.TempItemCreateRequest

	err = c.ShouldBindJSON(&item)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
			"error":   err.Error(),
		})
		return
	}

	tempItem := &schema.TempItemConvertedData{
		Id:        id,
		OriginURL: item.OriginURL,
		ShortPath: item.ShortPath,
		ExpSec:    item.ExpSec,
	}

	res, err := service.ConvertToTempItem(tempItem)

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
		"message": "성공적으로 변환되었습니다.",
		"result":  schema.NewItemResponse(res),
	})
	return
}

func convertToPermItem(c *gin.Context, err error, id int) {
	var item schema.PermItemCreateRequest

	err = c.ShouldBindJSON(&item)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
			"error":   err.Error(),
		})
		return
	}

	permItem := &schema.PermItemConvertedData{
		Id:        id,
		OriginURL: item.OriginURL,
		ShortPath: item.ShortPath,
	}

	res, err := service.ConvertToPermItem(permItem)

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
		"message": "성공적으로 변환되었습니다.",
		"result":  schema.NewItemResponse(res),
	})
	return
}

func DeleteItem(c *gin.Context) {

	reqId := c.Param("id")

	id, err := strconv.Atoi(reqId)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "잘못된 요청입니다.",
		})
		return
	}

	err = service.DeleteItem(id)

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
		"message": "성공적으로 삭제되었습니다.",
	})
}
