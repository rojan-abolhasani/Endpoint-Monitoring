package auth

import (
	"monitor/config"
	"time"

	"github.com/golang-jwt/jwt"
)

// creates a token sets the expiration date and the user id in the jwt token
func CreateToken(id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(config.TokenDuration).Unix(),
		"id":  id,
	})
	tokenString, err := token.SignedString(config.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
