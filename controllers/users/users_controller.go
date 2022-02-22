package users

import (
	"fmt"
	"net/http"

	"github.com/armuh16/book_user-api/domain/users"
	"github.com/armuh16/book_user-api/services"
	"github.com/gin-gonic/gin"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Print("Hallo world")
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO: Handle Error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	// TODO: Handler json error
	// 	return
	// }
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: Handle error
		fmt.Println(err)
		return
	}
	fmt.Println(user)
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO : Hadle user creating error
		return
	}

	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
