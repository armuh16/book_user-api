package app

import (
	"github.com/armuh16/book_user-api/controllers"
)

func mapUrls() {
	router.GET(relativePath: "/ping", controllers.Ping)
}
