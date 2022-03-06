package users

import (
	"github.com/armuh16/book_user-api/utils/errors"
)

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

func (user User) Save() *errors.RestErr {
	if userDB[user.Id] != nil {
		return errors.NewBadRequestError("user already exist")
	}
	return nil
}
