package repository

import (
	"database/sql"
	"github.com/hrabit64/shortlink/app/model"
	"github.com/hrabit64/shortlink/app/utils"
)

func CreateItem(db *sql.Tx, item model.Item) (int64, error) {

	query := `
        INSERT INTO ITEM (
            ITEM_ORIGIN_URL, ITEM_TYPE, ITEM_SHORT_PATH,
            ITEM_INIT_COUNT, ITEM_CURRENT_COUNT, ITEM_EXP_SEC, ITEM_CREATE_TIME
        ) VALUES (?, ?, ?, ?, ?, ?, ?)
    `
	result, err := db.Exec(query, item.OriginURL, item.Type, item.ShortPath, item.InitCount, item.CurrentCount, item.ExpSec, item.CreateTime)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func GetItemByID(db *sql.Tx, id int) (model.Item, error) {
	query := `SELECT ITEM_PK, ITEM_ORIGIN_URL, ITEM_TYPE, ITEM_SHORT_PATH, ITEM_INIT_COUNT, ITEM_CURRENT_COUNT, ITEM_EXP_SEC, ITEM_CREATE_TIME FROM ITEM WHERE ITEM_PK = ?`

	var item model.Item
	row := db.QueryRow(query, id)

	err := row.Scan(&item.ID, &item.OriginURL, &item.Type, &item.ShortPath, &item.InitCount, &item.CurrentCount, &item.ExpSec, &item.CreateTime)

	if err != nil {
		return item, err
	}

	return item, nil
}

func GetAllItems(db *sql.Tx, pageable *utils.Pageable) ([]model.Item, error) {

	query := `SELECT ITEM_PK, ITEM_ORIGIN_URL, ITEM_TYPE, ITEM_SHORT_PATH, ITEM_INIT_COUNT, ITEM_CURRENT_COUNT, ITEM_EXP_SEC, ITEM_CREATE_TIME FROM ITEM ORDER BY ITEM_PK DESC LIMIT ? OFFSET ?`

	rows, err := db.Query(query, pageable.Limit(), pageable.Offset())

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []model.Item

	for rows.Next() {
		var item model.Item
		err := rows.Scan(&item.ID, &item.OriginURL, &item.Type, &item.ShortPath, &item.InitCount, &item.CurrentCount, &item.ExpSec, &item.CreateTime)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func GetItemsCount(db *sql.Tx) (int, error) {
	query := `SELECT COUNT(*) FROM ITEM`

	var count int
	row := db.QueryRow(query)

	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func UpdateItem(db *sql.Tx, item model.Item) error {

	query := `
        UPDATE ITEM SET
            ITEM_ORIGIN_URL = ?, ITEM_TYPE = ?, ITEM_SHORT_PATH = ?, 
            ITEM_INIT_COUNT = ?, ITEM_CURRENT_COUNT = ?, ITEM_EXP_SEC = ?
        WHERE ITEM_PK = ?
    `
	_, err := db.Exec(query, item.OriginURL, item.Type, item.ShortPath, item.InitCount, item.CurrentCount, item.ExpSec, item.ID)
	return err
}

func DeleteItem(db *sql.Tx, id int) error {
	query := `DELETE FROM ITEM WHERE ITEM_PK = ?`

	_, err := db.Exec(query, id)

	return err
}

func GetItemByOriginURL(db *sql.Tx, originURL string) (model.Item, error) {
	query := `SELECT ITEM_PK, ITEM_ORIGIN_URL, ITEM_TYPE, ITEM_SHORT_PATH, ITEM_INIT_COUNT, ITEM_CURRENT_COUNT, ITEM_EXP_SEC, ITEM_CREATE_TIME FROM ITEM WHERE ITEM_ORIGIN_URL = ?`

	var item model.Item
	row := db.QueryRow(query, originURL)

	err := row.Scan(&item.ID, &item.OriginURL, &item.Type, &item.ShortPath, &item.InitCount, &item.CurrentCount, &item.ExpSec, &item.CreateTime)

	if err != nil {
		return item, err
	}

	return item, nil
}

func GetItemByShortPath(db *sql.Tx, shortPath string) (model.Item, error) {
	query := `SELECT ITEM_PK, ITEM_ORIGIN_URL, ITEM_TYPE, ITEM_SHORT_PATH, ITEM_INIT_COUNT, ITEM_CURRENT_COUNT, ITEM_EXP_SEC, ITEM_CREATE_TIME FROM ITEM WHERE ITEM_SHORT_PATH = ?`

	var item model.Item
	row := db.QueryRow(query, shortPath)

	err := row.Scan(&item.ID, &item.OriginURL, &item.Type, &item.ShortPath, &item.InitCount, &item.CurrentCount, &item.ExpSec, &item.CreateTime)

	if err != nil {
		return item, err
	}

	return item, nil
}
