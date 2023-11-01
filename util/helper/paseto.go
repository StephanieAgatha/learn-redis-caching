package helper

import (
	"github.com/gookit/slog"
	"github.com/o1egl/paseto/v2"
	"os"
	"strconv"
	"time"
)

func GeneratePaseto(email string, symetricKey string) string {
	key := []byte(symetricKey)
	//set token exp

	now := time.Now()
	expStr := os.Getenv("PASETO_EXP")
	exp, err := strconv.Atoi(expStr)
	if err != nil {
		return ""
	}
	expire := now.Add(time.Duration(exp) * time.Hour)

	jsonToken := paseto.JSONToken{
		Issuer:     "Soraa Go",
		Subject:    "Abrakadabra",
		Expiration: expire,
		IssuedAt:   now,
	}

	//custom claim
	jsonToken.Set("email", email)
	footer := "footer goes here"

	//encrypt
	token, err := paseto.Encrypt(key, jsonToken, footer)
	if err != nil {
		slog.Errorf("Failed to ecnrypt paseto %v", err.Error())
		return ""
	}
	return token
}

//func ParsePaseto(token string) (paseto.JSONToken, error) {
//	var jsonToken paseto.JSONToken
//	footer := ""
//
//	// Decrypt
//	err := paseto.Decrypt(token, symetricKey, &jsonToken, &footer)
//	if err != nil {
//		slog.Errorf("Failed to decrypt Paseto token: %v", err)
//		return paseto.JSONToken{}, err
//	}
//
//	return jsonToken, nil
//}


