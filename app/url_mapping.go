package app

import (
	pingcontroller "github.com/armuh16/book_user-api/controllers/ping"
	"github.com/armuh16/book_user-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", pingcontroller.Ping) // Path

	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
