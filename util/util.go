package util

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Arch-4ng3l/TaskManage/types"
	"github.com/golang-jwt/jwt/v5"
)

func AuthJWT(tokenString string, acc *types.Account) bool {

	token, err := ValidateJWT(tokenString)
	if err != nil {
		return false
	}
	if !token.Valid {
		return false
	}

	claims := token.Claims.(jwt.MapClaims)

	name := claims["username"]
	email := claims["email"]

	if email != acc.Email || name != acc.Username {
		return false
	}
	return true

}

func CreateJWT(acc *types.Account) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"username":  acc.Username,
		"email":     acc.Email,
	}
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

}

func GetCookies(r *http.Request) ([3]string, error) {
	var arr [3]string
	name, err := r.Cookie("username")
	if err != nil {
		return arr, err
	}
	email, err := r.Cookie("email")
	if err != nil {
		return arr, err
	}
	token, err := r.Cookie("jwt-token")

	if err != nil {
		return arr, err
	}

	arr[0] = name.Value
	arr[1] = email.Value
	arr[2] = token.Value
	return arr, nil

}
