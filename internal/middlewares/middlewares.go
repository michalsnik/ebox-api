package middlewares

import "github.com/gin-gonic/gin"

type AuthService interface {
	ParseToken(tokenString string) (int, error)
}

type middlewares struct {
	AuthRequired gin.HandlerFunc
}

func Initialize(authService AuthService) *middlewares {
	return &middlewares{
		AuthRequired: NewAuthRequiredMiddleware(authService),
	}
}