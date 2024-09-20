package jwt

import (
	"fmt"
	"log"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func Parsejwt(token string) (string, error) {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("secret key not set in environment")
	}

	claims := jwt.MapClaims{}
	_,err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}
	return claims["email"].(string), nil
}