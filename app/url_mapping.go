package app

import (
	"bookapi/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
