package services

import (
	"net/http"

	"github.com/armuh16/book_user-api/domain/users"
	"github.com/armuh16/book_user-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// return nil, nil

	// return &user, nil

	return &user, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}

func GetUser() {

}

func FindUser() {

}
