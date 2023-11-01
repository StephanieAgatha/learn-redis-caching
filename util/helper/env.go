package helper

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Loadenv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Failed to load .env file %v", err.Error())
	}
	return nil
}

