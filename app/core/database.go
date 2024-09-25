package core

import (
	"database/sql"
	"fmt"
	"io"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnect() (*sql.DB, error) {

	conn, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		return nil, err
	}

	return conn, err
}

func InitDB(sqlFilePath string) error {
	conn, err := GetConnect()

	if err != nil {
		return err
	}

	defer conn.Close()

	// SQL 파일을 읽어옴
	initSql, err := readSql(sqlFilePath)

	if err != nil {
		return err
	}

	// 읽은 SQL 구문 실행
	_, err = conn.Exec(initSql)

	if err != nil {
		return err
	}

	//check user table, if not exist, create user
	cntRow := conn.QueryRow("SELECT COUNT(*) FROM USER")

	var count int
	err = cntRow.Scan(&count)

	if err != nil {
		return err
	}

	if count == 0 {
		_, err = conn.Exec("INSERT INTO USER (USER_ID, USER_PW) VALUES ('admin', 'admin')")
		if err != nil {
			return err
		}
	}

	fmt.Println("Database initialized successfully!")

	return nil
}

func readSql(filePath string) (string, error) {
	sqlFile, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer sqlFile.Close()

	// 파일 내용 읽기
	sqlContent, err := io.ReadAll(sqlFile)
	if err != nil {
		return "", err
	}

	return string(sqlContent), nil
}
