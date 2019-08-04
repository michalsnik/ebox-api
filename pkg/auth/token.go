package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

const (
	hmacSecret = "asdqlwepoque1092309uhaksdan1k2m0129"
)

func tokenParser(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(hmacSecret), nil
}

// Returns userID if token is valid and contains the right claim
func ParseToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, tokenParser)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return 0, errors.New("meh")
	}

	userID, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func CreateToken(userID int) (string, error) {
	claims := &jwt.StandardClaims{
		Issuer: strconv.Itoa(userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}