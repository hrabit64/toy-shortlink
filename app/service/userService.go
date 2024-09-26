package service

import (
	"github.com/hrabit64/shortlink/app/core"
	"github.com/hrabit64/shortlink/app/model"
	"github.com/hrabit64/shortlink/app/repository"
	"github.com/hrabit64/shortlink/app/schema"
	"github.com/hrabit64/shortlink/app/utils"
)

func GetUserById(id string) (model.User, error) {
	conn, err := core.GetConnect()

	if err != nil {
		return model.User{}, err
	}

	defer conn.Close()

	return repository.GetUserById(conn, id)
}

func UpdateUser(req schema.UserUpdateRequest) (int64, error) {
	conn, err := core.GetConnect()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	hashPw, err := utils.HashPassword(req.Pw)

	if err != nil {
		return 0, err
	}

	user := model.User{
		Id: req.Id,
		Pw: hashPw,
	}
	res, err := repository.UpdateUser(conn, user)

	if err != nil {
		return 0, err
	}

	return res, nil
}
