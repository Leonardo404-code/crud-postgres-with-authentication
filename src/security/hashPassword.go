package security

import "golang.org/x/crypto/bcrypt"

// HashPassword Generate a Hash for user password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}
