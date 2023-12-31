package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateNewAccessToken()(string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	// Expiration time
	expirationTimeInMins := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expirationTimeInMins)).Unix()

	token := jwt.NewWithClaims(jwt.SigninMethodHS256, claims)

	newToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", nil
	}

	return newToken, nil
}