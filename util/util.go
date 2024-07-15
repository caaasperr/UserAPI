package util

import (
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func InitializeEnv() {
	err := godotenv.Load()
	if err != nil {
		Logger.Error("Error loading .env file")
	}
}

func Pass2Hash(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "0"
	}
	return string(hashed)
}

func ComparePassword(password string, hash string) bool {
	er := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if er != nil {
		return false
	}
	return true
}
