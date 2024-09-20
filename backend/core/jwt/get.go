package jwt

import (
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)


func GenerateJWT(email string) (string, error) {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("secret key not set in environment")
	}

	token,err :=jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":email,
			"exp":time.Now().Add(time.Hour * 100).Unix(),
		}).SignedString([]byte(secretKey))
		if err!=nil{
			return "",err
		}
		return token,nil}