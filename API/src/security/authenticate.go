package security

import (
	"example.com/m/v2/API/src/configuration"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(userId int64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString(configuration.SecretKey)
}
