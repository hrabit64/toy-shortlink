package schema

import (
	"github.com/hrabit64/shortlink/app/model"
	"github.com/hrabit64/shortlink/app/utils"
	"time"
)

type ItemPage struct {
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	LastPage int             `json:"last_page"`
	Total    int             `json:"total"`
	Items    []*ItemResponse `json:"items"`
}

func NewItemPage(pageable utils.Pageable, total int, items []model.Item) *ItemPage {

	lastPage := total / pageable.PageSize

	return &ItemPage{
		Page:     pageable.Page,
		PageSize: pageable.PageSize,
		LastPage: lastPage,
		Total:    len(items),
		Items:    ConvertItemsToResponses(items),
	}
}

type ItemResponse struct {
	ID           int       `json:"id"`
	OriginURL    string    `json:"origin_url"`
	ShortPath    string    `json:"short_path"`
	CurrentCount int64     `json:"current_count"`
	ExpSec       int64     `json:"exp_sec"`
	Type         string    `json:"type"`
	InitCount    int64     `json:"init_count"`
	CreateTime   time.Time `json:"create_time"`
}

func NewItemResponse(item model.Item) *ItemResponse {

	return &ItemResponse{
		ID:           item.ID,
		OriginURL:    item.OriginURL,
		ShortPath:    item.ShortPath,
		CurrentCount: item.GetCurrentCount(),
		ExpSec:       item.GetExpSec(),
		InitCount:    item.GetInitCount(),
		CreateTime:   item.CreateTime,
		Type:         item.Type,
	}
}

func ConvertItemsToResponses(items []model.Item) []*ItemResponse {
	var responses []*ItemResponse
	for _, item := range items {
		response := NewItemResponse(item)
		responses = append(responses, response) // 포인터 역참조
	}
	return responses
}

type PermItemCreateRequest struct {
	OriginURL string `json:"origin_url" binding:"required,url,max=2083"`
	ShortPath string `json:"short_path" binding:"required,min=1,max=50"`
}

type TempItemCreateRequest struct {
	OriginURL string `json:"origin_url" binding:"required,url,max=2083"`
	ShortPath string `json:"short_path" binding:"required,min=1,max=50"`
	ExpSec    int64  `json:"exp_sec" binding:"required,min=60,max=86400"`
}

type CountItemCreateRequest struct {
	OriginURL string `json:"origin_url" binding:"required,url,max=2083"`
	ShortPath string `json:"short_path" binding:"required,min=1,max=50"`
	InitCount int64  `json:"init_count" binding:"required,min=1,max=100"`
}

type PermItemUpdateRequest struct {
	OriginURL string `json:"origin_url" binding:"required,url,max=2083"`
	ShortPath string `json:"short_path" binding:"required,min=1,max=50"`
}

type TempItemUpdateRequest struct {
	OriginURL string `json:"origin_url" binding:"required,url,max=2083"`
	ShortPath string `json:"short_path" binding:"required,min=1,max=50"`
	ExpSec    int64  `json:"exp_sec" binding:"required,min=0,max=86400"`
}

type CountItemUpdateRequest struct {
	OriginURL string `json:"origin_url" binding:"required,url,max=2083"`
	ShortPath string `json:"short_path" binding:"required,min=0,max=50"`
	InitCount int64  `json:"init_count" binding:"required,min=0,max=100"`
}

type PermItemCreateData struct {
	OriginURL string
	ShortPath string
}

type TempItemCreateData struct {
	OriginURL string
	ShortPath string
	ExpSec    int64
}

type CountItemCreateData struct {
	OriginURL string
	ShortPath string
	InitCount int64
}

type PermItemUpdateData struct {
	Id        int
	OriginURL string
	ShortPath string
}

type TempItemUpdateData struct {
	Id        int
	OriginURL string
	ShortPath string
	ExpSec    int64
}

type CountItemUpdateData struct {
	Id        int
	OriginURL string
	ShortPath string
	InitCount int64
}

type PermItemConvertedData struct {
	Id        int
	OriginURL string
	ShortPath string
}

type TempItemConvertedData struct {
	Id        int
	OriginURL string
	ShortPath string
	ExpSec    int64
}

type CountItemConvertedData struct {
	Id        int
	OriginURL string
	ShortPath string
	InitCount int64
}
