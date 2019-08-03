package auth

import (
	"ebox-api/internal/db"
	"ebox-api/pkg/users"
	"github.com/gin-gonic/gin"
)

func Register (router *gin.RouterGroup, db *db.DB) {
	usersRepository := users.NewUsersRepository(db)
	authSvc := NewAuthService(usersRepository)
	handlers := NewAuthHandlers(authSvc)

	router.POST("/auth/sign-in", handlers.SignIn)
}
