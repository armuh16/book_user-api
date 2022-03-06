package users

import (
	"github.com/armuh16/book_user-api/utils/errors"
)

// connect to db

var (
	userDB = make(map[int64]*User)
)

// func something() {
// 	user := User{}

// 	if err := user.Get(); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(user.FirstName)
// }

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError("user id not found")
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError("email already registered")
		}
		return errors.NewBadRequestError("user already exist")
	}
	userDB[user.Id] = user
	return nil
}
