package dao

import (
	"fmt"
	"net/http"
	"users/errors"
	"users/model"
)

var (
	users = map[string]*model.User{
		"test": {UserName: "javad", Email: "javad.8558@gmail.com", Password: "easypeasy"},
	}
)

func GetUser(username string) (*model.User, *errors.AppError) {
	if user := users[username]; user != nil {
		return user, nil
	}

	return nil, &errors.AppError{
		Message:    fmt.Sprintf("User %v was not found", username),
		StatusCode: http.StatusNotFound,
		Status:     "not_found",
	}
}
