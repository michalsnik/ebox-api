package auth

import (
	"ebox-api/pkg/users"
	"github.com/dgrijalva/jwt-go"
)

const (
	hmacSecret = "asdqlwepoque1092309uhaksdan1k2m0129"
)

type AuthService struct {
	usersRepository *users.UsersRepository
}

func NewAuthService(usersRepository *users.UsersRepository) *AuthService {
	return &AuthService{
		usersRepository: usersRepository,
	}
}

func (svc *AuthService) AuthenticateUser (email string, password string) (string, error) {
	userID, err := svc.usersRepository.ValidateUser(email, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
	})

	tokenString, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
