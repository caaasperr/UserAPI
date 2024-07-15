package util

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateAccessToken(username string) (string, error) {
	secretKey := os.Getenv("Secret-Key")
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["Username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	t, err := refreshToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func CreateRefreshToken(username string) (string, error) {
	secretKey := os.Getenv("Secret-Key")
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["Username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix()
	t, err := refreshToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func VerifyAccessToken(accessToken string) bool {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", nil
		}
		return []byte(os.Getenv("Secret-Key")), nil
	})
	if err != nil || !token.Valid {
		return false
	}
	return true
}

func GetUsernameFromToken(accessToken string) string {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", nil
		}
		return []byte(os.Getenv("Secret-Key")), nil
	})
	if err != nil || !token.Valid {
		return ""
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return ""
	}
	username := claims["Username"].(string)
	return username
}
