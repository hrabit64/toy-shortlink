package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/hrabit64/shortlink/app/userErrors"
	"strconv"
)

type Pageable struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// Offset returns the offset of the current page
func (p *Pageable) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// Limit returns the limit of the current page
func (p *Pageable) Limit() int {
	return p.PageSize
}

func NewPageable(c *gin.Context) (*Pageable, error) {

	page, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		return nil, &userErrors.BusinessError{Message: "page 값이 잘못되었습니다.", Status: 400}
	}

	pageSize, err := strconv.Atoi(c.Query("size"))

	if err != nil {
		return nil, &userErrors.BusinessError{Message: "size 값이 잘못되었습니다.", Status: 400}
	}

	if pageSize >= 100 {
		return nil, &userErrors.BusinessError{Message: "size 값이 너무 큽니다.", Status: 400}
	}

	pageable := Pageable{
		Page:     page,
		PageSize: pageSize,
	}
	return &pageable, nil
}
