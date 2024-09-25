package service

import (
	"database/sql"
	"errors"
	"github.com/hrabit64/shortlink/app/core"
	"github.com/hrabit64/shortlink/app/model"
	"github.com/hrabit64/shortlink/app/repository"
	"github.com/hrabit64/shortlink/app/schema"
	"github.com/hrabit64/shortlink/app/userErrors"
	"github.com/hrabit64/shortlink/app/utils"
	"log"
	"time"
)

// 아이템을 ID로 조회한다. 이때 아이템이 존재하지 않는 경우 BusinessError 에러를 반환한다.
func GetItemById(id int) (model.Item, error) {

	conn, err := core.GetConnect()

	if err != nil {
		return model.Item{}, err
	}

	defer conn.Close()

	tx, err := conn.Begin()

	if err != nil {
		return model.Item{}, err
	}

	item, err := repository.GetItemByID(tx, id)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return model.Item{}, &userErrors.BusinessError{Message: "해당 아이템을 찾을 수 없습니다.", Status: 404}
		}

		return model.Item{}, err
	}

	return item, nil
}

func CreatePermItem(data schema.PermItemCreateData) (int64, error) {

	item := model.Item{
		OriginURL:  data.OriginURL,
		ShortPath:  data.ShortPath,
		Type:       "perm",
		CreateTime: time.Now(),
	}

	return CreateItem(item)
}

func CreateTempItem(data schema.TempItemCreateData) (int64, error) {

	item := model.Item{
		OriginURL: data.OriginURL,
		ShortPath: data.ShortPath,
		Type:      "temp",
		ExpSec: sql.NullInt64{
			Int64: data.ExpSec,
			Valid: true,
		},
		CreateTime: time.Now(),
	}

	return CreateItem(item)
}

func CreateCountItem(data schema.CountItemCreateData) (int64, error) {

	item := model.Item{
		OriginURL: data.OriginURL,
		ShortPath: data.ShortPath,
		Type:      "count",
		InitCount: sql.NullInt64{
			Int64: data.InitCount,
			Valid: true,
		},
		CurrentCount: sql.NullInt64{
			Int64: data.InitCount,
			Valid: true,
		},
		CreateTime: time.Now(),
	}

	return CreateItem(item)
}

// 아이템을 생성한다. 이때 원본 URL과 단축 URL이 이미 존재하는 경우 BusinessError 에러를 반환한다.
func CreateItem(item model.Item) (int64, error) {

	conn, err := core.GetConnect()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	tx, err := conn.Begin()

	if err != nil {
		return 0, nil
	}

	defer func() {
		if err != nil {
			log.Println(err)
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = repository.GetItemByOriginURL(tx, item.OriginURL)

	if err == nil {
		return 0, &userErrors.BusinessError{Message: "해당 원본 URL은 이미 존재합니다.", Status: 400}
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	_, err = repository.GetItemByShortPath(tx, item.ShortPath)

	if err == nil {
		return 0, &userErrors.BusinessError{Message: "해당 단축 URL은 이미 존재합니다.", Status: 400}
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	result, err := repository.CreateItem(tx, item)

	if err != nil {
		return 0, err
	}

	return result, nil
}

func UpdatePermItem(data schema.PermItemUpdateData) error {

	item := model.Item{
		ID:        data.Id,
		OriginURL: data.OriginURL,
		ShortPath: data.ShortPath,
	}

	return UpdateItem(item)
}

func UpdateTempItem(data schema.TempItemUpdateData) error {

	item := model.Item{
		ID:        data.Id,
		OriginURL: data.OriginURL,
		ShortPath: data.ShortPath,
		ExpSec: sql.NullInt64{
			Int64: data.ExpSec,
			Valid: true,
		},
	}

	return UpdateItem(item)
}

func UpdateCountItem(data schema.CountItemUpdateData) error {

	item := model.Item{
		ID:        data.Id,
		OriginURL: data.OriginURL,
		ShortPath: data.ShortPath,
		InitCount: sql.NullInt64{
			Int64: data.InitCount,
			Valid: true,
		},
		CurrentCount: sql.NullInt64{
			Int64: data.InitCount,
			Valid: true,
		},
	}

	return UpdateItem(item)
}

// 모든 아이템을 조회한다. 이때 아이템이 존재하지 않는 경우 BusinessError 에러를 반환한다.
func GetAllItems(pageable *utils.Pageable) (schema.ItemPage, error) {

	conn, err := core.GetConnect()

	if err != nil {
		return schema.ItemPage{}, err
	}

	defer conn.Close()

	tx, err := conn.Begin()

	if err != nil {
		return schema.ItemPage{}, err
	}

	items, err := repository.GetAllItems(tx, pageable)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return schema.ItemPage{}, &userErrors.BusinessError{Message: "아이템을 찾을 수 없습니다.", Status: 404}
		}

		return schema.ItemPage{}, err
	}

	total, err := repository.GetItemsCount(tx)

	if err != nil {
		return schema.ItemPage{}, err
	}

	return *schema.NewItemPage(*pageable, total, items), nil
}

// 아이템을 수정한다. 아이템이 존재하지 않는 경우 BusinessError 에러를 반환한다.
func UpdateItem(item model.Item) error {

	conn, err := core.GetConnect()

	if err != nil {
		return err
	}

	defer conn.Close()

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// 아이템이 존재하는지 확인
	_, err = repository.GetItemByID(tx, item.ID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &userErrors.BusinessError{Message: "해당 아이템을 찾을 수 없습니다.", Status: 404}
		}
		return err
	}

	err = repository.UpdateItem(tx, item)

	if err != nil {
		return err
	}

	return nil
}

func DeleteItem(id int) error {

	conn, err := core.GetConnect()

	if err != nil {
		return err
	}

	defer conn.Close()

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = repository.DeleteItem(tx, id)

	if err != nil {
		return err
	}

	return nil
}

func ConvertToPermItem(data *schema.PermItemConvertedData) (model.Item, error) {

	conn, err := core.GetConnect()

	if err != nil {
		return model.Item{}, err
	}

	defer conn.Close()

	tx, err := conn.Begin()

	if err != nil {
		return model.Item{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	item, err := repository.GetItemByID(tx, data.Id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Item{}, &userErrors.BusinessError{Message: "해당 아이템을 찾을 수 없습니다.", Status: 404}
		}
		return model.Item{}, err
	}

	if item.Type == "perm" {
		return model.Item{}, &userErrors.BusinessError{Message: "이미 영구 URL 입니다.", Status: 400}
	}

	item.Type = "perm"
	item.InitCount = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
	item.CurrentCount = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
	item.ExpSec = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}

	err = repository.UpdateItem(tx, item)

	if err != nil {
		return model.Item{}, err
	}

	return item, nil

}

func ConvertToTempItem(data *schema.TempItemConvertedData) (model.Item, error) {

	conn, err := core.GetConnect()

	if err != nil {
		return model.Item{}, err
	}

	defer conn.Close()

	tx, err := conn.Begin()

	if err != nil {
		return model.Item{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	item, err := repository.GetItemByID(tx, data.Id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Item{}, &userErrors.BusinessError{Message: "해당 아이템을 찾을 수 없습니다.", Status: 404}
		}
		return model.Item{}, err
	}

	if item.Type == "temp" {
		return model.Item{}, &userErrors.BusinessError{Message: "이미 임시 URL 입니다.", Status: 400}
	}

	item.Type = "temp"
	item.InitCount = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
	item.CurrentCount = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}

	err = repository.UpdateItem(tx, item)

	if err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func ConvertToCountItem(data *schema.CountItemConvertedData) (model.Item, error) {

	conn, err := core.GetConnect()

	if err != nil {
		return model.Item{}, err
	}

	defer conn.Close()

	tx, err := conn.Begin()

	if err != nil {
		return model.Item{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	item, err := repository.GetItemByID(tx, data.Id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Item{}, &userErrors.BusinessError{Message: "해당 아이템을 찾을 수 없습니다.", Status: 404}
		}
		return model.Item{}, err
	}

	if item.Type == "count" {
		return model.Item{}, &userErrors.BusinessError{Message: "이미 카운트 URL 입니다.", Status: 400}
	}

	item.Type = "count"
	item.ExpSec = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}

	err = repository.UpdateItem(tx, item)

	if err != nil {
		return model.Item{}, err
	}

	return item, nil
}

// 단축 URL로 아이템을 조회하고, URL 타입에 따라 만료 등의 처리를 수행한다.
func LookUpItemByShortPath(shortPath string) (model.Item, error) {

	conn, err := core.GetConnect()

	if err != nil {
		return model.Item{}, err
	}

	defer conn.Close()

	tx, err := conn.Begin()

	if err != nil {
		return model.Item{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	item, err := repository.GetItemByShortPath(tx, shortPath)

	if err != nil {
		return model.Item{}, err
	}

	//'temp', 'perm', 'count'
	return checkItemIsValid(item, err, tx)
}

func checkItemIsValid(item model.Item, err error, tx *sql.Tx) (model.Item, error) {
	if item.Type == "perm" {
		return item, nil

	} else if item.Type == "temp" {
		return processTempType(item, err, tx)
	} else if item.Type == "count" {
		return processCountType(item, err, tx)
	} else {
		return model.Item{}, &userErrors.BusinessError{Message: "알 수 없는 URL 타입입니다.", Status: 500}
	}
}

// 카운트 타입의 아이템을 처리한다.
func processCountType(item model.Item, err error, tx *sql.Tx) (model.Item, error) {
	if item.CurrentCount.Valid {
		if item.CurrentCount.Int64 <= 0 {
			return model.Item{}, &userErrors.BusinessError{Message: "해당 URL을 찾을 수 없습니다.", Status: 404}
		}

		item.CurrentCount.Int64--
		err = repository.UpdateItem(tx, item)
		if err != nil {
			return model.Item{}, err
		}

		return item, nil
	}

	return item, &userErrors.BusinessError{
		Message: "데이터가 올바르지 않습니다.",
		Status:  500,
	}
}

// 임시 타입의 아이템을 처리한다.
func processTempType(item model.Item, err error, tx *sql.Tx) (model.Item, error) {
	if item.IsExpired() {
		return model.Item{}, &userErrors.BusinessError{Message: "해당 URL을 찾을 수 없습니다.", Status: 404}
	}
	return item, nil
}
