package repository

import (
	"database/sql"
	"github.com/hrabit64/shortlink/app/model"
)

func GetUserById(db *sql.DB, id string) (model.User, error) {
	query := `SELECT USER_PK, USER_ID, USER_PW FROM USER WHERE USER_ID = ?`

	var user model.User
	row := db.QueryRow(query, id)

	err := row.Scan(&user.Pk, &user.Id, &user.Pw)

	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(db *sql.DB, user model.User) (int64, error) {
	query := `
		INSERT INTO USER (
			USER_ID, USER_PW
		) VALUES (?, ?)
	`
	result, err := db.Exec(query, user.Id, user.Pw)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func UpdateUser(db *sql.DB, user model.User) (int64, error) {
	query := `
		UPDATE USER SET USER_PW = ?, USER_ID = ? WHERE USER_ID = ?
	`
	result, err := db.Exec(query, user.Pw, user.Id, user.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
