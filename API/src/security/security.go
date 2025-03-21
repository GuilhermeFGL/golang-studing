package security

import "golang.org/x/crypto/bcrypt"

// Hash transform a password to hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPasswordHash compare a password to a hash
func CheckPasswordHash(hash, text string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(text)) == nil
}
