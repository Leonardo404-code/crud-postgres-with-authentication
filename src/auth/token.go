package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// CreateToken return a token with user id and expiration
func CreateToken(id uint) (string, error) {
	secret := []byte(os.Getenv("TOKEN_SECRET"))

	permission := jwt.MapClaims{}
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permission["id"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)

	return token.SignedString(secret)
}

// ExtractUserID return user ID in JWT request
func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)

	token, erro := jwt.Parse(tokenString, returnKeyVerification)

	if erro != nil {
		return 0, erro
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, erro := strconv.ParseUint(
			fmt.Sprintf("%.0f", permissions["id"]), 10, 64,
		)

		if erro != nil {
			return 0, erro
		}

		return userID, nil
	}

	return 0, errors.New("token inválido")
}

// extractToken return token in request
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// returnKeyVerification return verification if JWT is valid
func returnKeyVerification(token *jwt.Token) (interface{}, error) {
	secret := []byte(os.Getenv("TOKEN_SECRET"))
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Metodo de assinatura inesperado! %v", token.Header["alg"])
	}

	return secret, nil
}
