package users

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/armuh16/book_user-api/domain/users"
	"github.com/gin-gonic/gin"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO: Handle Error
		return
	}
	fmt.Println(string(bytes))
	fmt.Println(err)
	c.String(http.StatusNotImplemented, "implement me")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
