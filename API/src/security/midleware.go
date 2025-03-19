package security

import (
	"errors"
	"example.com/m/v2/API/src/configuration"
	"example.com/m/v2/API/src/util/httpresponse"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Logger log request
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next(w, r)
	}
}

// Authorize authorize user to make request
func Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := validateToken(r); err != nil {
			httpresponse.Error(w, http.StatusUnauthorized, err.Error())
			return
		}
		next(w, r)
	}
}

// ExtractUserIdFromToken return user id from token
func ExtractUserIdFromToken(r *http.Request) (uint64, error) {
	headerToken := extractToken(r)
	token, err := jwt.Parse(headerToken, getSecretKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprint(permissions["user_id"]), 10, 64)
		if err == nil {
			return userId, nil
		}
	}

	return 0, errors.New("invalid token")
}

func validateToken(r *http.Request) error {
	headerToken := extractToken(r)
	token, err := jwt.Parse(headerToken, getSecretKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	} else {
		return errors.New("invalid token")
	}
}

func extractToken(r *http.Request) string {
	return strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
}

func getSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("wrong signing algorithim %v", token.Header["alg"])
	}

	return configuration.SecretKey, nil
}
