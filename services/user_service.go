package services

import (
	"users/dao"
	"users/errors"
	"users/model"
)

func GetUser(username string) (*model.User, *errors.AppError) {
	return dao.GetUser(username)
}
