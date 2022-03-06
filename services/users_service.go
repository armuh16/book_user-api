package services

import (
	"github.com/armuh16/book_user-api/domain/users"
	"github.com/armuh16/book_user-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
	// return &user, &errors.RestErr{
	// 	Status: http.StatusInternalServerError,
	// }
}

func GetUser() {

}

func FindUser() {

}
