package model

type User struct {
	Pk int    `json:"user_pk"` // USER_PK
	Id string `json:"user_id"` // USER_ID
	Pw string `json:"user_pw"` // USER_PW
}
