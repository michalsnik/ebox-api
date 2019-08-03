package users

import (
	"ebox-api/internal/db"
	"github.com/gin-gonic/gin"
)

func Register (router *gin.RouterGroup, db *db.DB) {
	usersRepository := NewUsersRepository(db)
	usersSvc := NewUsersService(usersRepository)
	handlers := NewUsersHandlers(usersSvc)

	router.POST("/users", handlers.PostUser)
}