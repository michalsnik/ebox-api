package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	headerPrefix = "Bearer "
)

var (
	ErrWrongAuthorizationHeader = errors.New("wrong authorization header")
)

func getTokenFromHeader(header string) (string, error) {
	isValidHeader := strings.HasPrefix(header, headerPrefix)
	if !isValidHeader {
		return "", ErrWrongAuthorizationHeader
	}

	token := strings.Replace(header, headerPrefix, "", 1)

	return token, nil
}

func NewAuthRequiredMiddleware(service AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token, err := getTokenFromHeader(authHeader)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		issuer, err := service.ParseToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Set("userID", issuer)
		c.Next()
	}
}
