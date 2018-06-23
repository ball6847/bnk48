package utils

import (
	"time"

	flag "github.com/ball6847/bnk48/flag"
	p "github.com/ball6847/bnk48/payload"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken from payload
func GenerateToken(data *p.Signup, secret string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(*flag.Secret))
}
