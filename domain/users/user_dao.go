package users

import (
	"fmt"

	"github.com/armuh16/book_user-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func something() {
	user := User{}

	if err := user.Get(); err != nil {
		fmt.Println(err)
		return
	}
}

func (user User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Printf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user User) Save() *errors.RestErr {
	return nil
}
