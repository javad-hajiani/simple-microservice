package dao

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserError(t *testing.T) {
	user, err := GetUser("testkla")

	assert.NotNil(t, err, "username with testkla is not accepted")

	assert.Nil(t, user, "user should be nil in case of error")

	assert.Equal(t, http.StatusNotFound, err.StatusCode, "expecting 404 status code")
	//if user != nil {
	//t.Error("username with testkla is not accepted")
	//}

	//if err == nil {
	//t.Error("expecting error for username testkla")
	//}

	//if err.StatusCode != http.StatusNotFound {
	//t.Error("expecting 404 status code")
	//}
}

func TestGetUserNotError(t *testing.T) {
	user, err := GetUser("test")

	if user == nil {
		t.Error("username with test is expected")
	}

	if err != nil {
		t.Error("Not expecting error with username test")
	}
}
