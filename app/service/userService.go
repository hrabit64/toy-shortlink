package service

import (
	"github.com/hrabit64/shortlink/app/core"
	"github.com/hrabit64/shortlink/app/model"
	"github.com/hrabit64/shortlink/app/repository"
	"github.com/hrabit64/shortlink/app/schema"
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

	user := model.User{
		Id: req.Id,
		Pw: req.Pw,
	}
	res, err := repository.UpdateUser(conn, user)

	if err != nil {
		return 0, err
	}
	
	return res, nil
}
