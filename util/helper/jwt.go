package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"redishop/model"
	"time"
)

func GenerateJWT(user model.UserCredentials, secretjwt string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issued_at": time.Now(),
		"exp_at":    time.Now().Add(4 * time.Hour),
		"email":     user.Email,
	})

	tokenstr, err := token.SignedString(secretjwt)
	if err != nil {
		return "", fmt.Errorf("Token is invalid / expired")
	}

	return tokenstr, nil
}

// parse jwt
func ParseJWT(tknHeader string, secretjwt string) (*jwt.Token, error) {
	return jwt.Parse(tknHeader, func(token *jwt.Token) (interface{}, error) {
		return secretjwt, nil
	})
}
