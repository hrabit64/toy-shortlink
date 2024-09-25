package model

import (
	"database/sql"
	"time"
)

type Item struct {
	ID           int           `json:"id"`                      // ITEM_PK
	OriginURL    string        `json:"origin_url"`              // ITEM_ORIGIN_URL
	Type         string        `json:"type"`                    // ITEM_TYPE ('temp', 'perm', 'count')
	ShortPath    string        `json:"short_path"`              // ITEM_SHORT_PATH
	InitCount    sql.NullInt64 `json:"init_count,omitempty"`    // ITEM_INIT_COUNT
	CurrentCount sql.NullInt64 `json:"current_count,omitempty"` // ITEM_CURRENT_COUNT
	ExpSec       sql.NullInt64 `json:"exp_sec,omitempty"`       // ITEM_EXP_SEC
	CreateTime   time.Time     `json:"create_time"`             // ITEM_CREATE_TIME
}

func (item *Item) IsExpired() bool {

	// perm is always valid
	if item.Type == "perm" {
		return false
	}

	// count is valid until current count is 0
	if item.Type == "count" {
		if item.CurrentCount.Valid {
			return item.CurrentCount.Int64 <= 0
		}
		return false
	}

	// temp is valid until expiration time
	if item.Type == "temp" {
		if item.ExpSec.Valid {
			return time.Now().After(item.CreateTime.Add(time.Duration(item.ExpSec.Int64) * time.Second))
		}
		return false

	}

	return true
}

func (item *Item) GetCurrentCount() int64 {
	if item.CurrentCount.Valid {
		return item.CurrentCount.Int64
	}
	return 0
}

func (item *Item) GetExpSec() int64 {
	if item.ExpSec.Valid {
		return item.ExpSec.Int64
	}
	return 0
}

func (item *Item) GetInitCount() int64 {
	if item.InitCount.Valid {
		return item.InitCount.Int64
	}
	return 0
}
