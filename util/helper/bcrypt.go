package helper

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fmt.Errorf("Failed to hash")
	}

	return string(hashedPass), nil
}

func ComparePassword(hashpassword, password string) error {
	hashedPass := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	return hashedPass
}


