package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers (c *gin.Context) {
	users := []User{
		{
			Id: 1,
			Name: "Jon",
			Email: "jon.snow@gmail.com",
		},
		{
			Id: 2,
			Name: "Emma",
			Email: "emma.watson@gmail.com",
		},
	}

	c.JSON(http.StatusOK, users)
}