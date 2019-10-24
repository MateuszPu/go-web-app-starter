package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func Get() gin.HandlerFunc {
	return func(context *gin.Context) {
		name := context.Query("name")
		lastName := context.Query("lastName")
		age, _ := strconv.Atoi(context.DefaultQuery("age", "0"))
		context.JSON(http.StatusOK, User{name, lastName, age})
	}
}