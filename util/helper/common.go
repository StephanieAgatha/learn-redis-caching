package helper

import (
	"encoding/json"
	"fmt"
	"redishop/model"
	"regexp"
	"unicode"
)

func IsEmailValid(user model.UserCredentials) error {
	// Check if email is valid (e.g., gmail.com)

	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$` //change to tag
	validEmail := regexp.MustCompile(emailPattern)
	if !validEmail.MatchString(user.Email) {
		return fmt.Errorf("Invalid Email Format")
	}
	return nil
}

/*
---------PASSWORD REQUIREMENT AREA--------------
*/

func PasswordContainsUppercase(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func PasswordContainsSpecialChar(s string) bool {
	// Regular expression to match any special character
	re := regexp.MustCompile(`[!@#$%^&*()_+=\[{\]};:'",<.>/?]`) //change to tag
	return re.MatchString(s)
}

func PasswordConstainsOneNumber(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func JSONALLProductParse(val []byte) []model.Products {
	var products []model.Products

	if err := json.Unmarshal(val, &products); err != nil {
		fmt.Printf("Failed to parse JSON %v", err.Error())
		return nil
	}
	return products
}
