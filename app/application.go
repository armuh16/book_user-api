package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	fmt.Println("server starting...")
	mapUrls()
	router.Run(":8081") //address
}
