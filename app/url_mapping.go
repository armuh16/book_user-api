package app

import (
	"github.com/armuh16/book_user-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping) // Path

	router.GET("/users/:user_id", controllers.GetUser)
	router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", controllers.CreateUser)
}
