package services

import (
	"github.com/armuh16/book_user-api/domain/users"
	"github.com/armuh16/book_user-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}

func GetUser() {

}

func FindUser() {

}
