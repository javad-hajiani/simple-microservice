package controllers

import (
	"encoding/json"
	"net/http"
	"users/errors"
	"users/services"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	username := request.URL.Query().Get("username")

	if username == "" {
		apiErr := &errors.AppError{
			Message:    "username not provaded",
			StatusCode: http.StatusBadRequest,
			Status:     "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)

		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(username)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	response.Write(jsonValue)
}
