package services

import (
	"GolangBackend/config"
	"GolangBackend/helper"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserJWT struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Roles string `json:"roles"`
	jwt.RegisteredClaims
}

func SignJWT(data *UserJWT) (string, error) {
	JWT_SECRET := config.GetEnv("JWT_SECRET", "JWT_SECRET")
	expAt := time.Now().Add(24 * time.Hour)
	signData := &UserJWT{
		Name:  data.Name,
		Email: data.Email,
		Phone: data.Phone,
		Roles: data.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expAt),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, signData)
	return jwtToken.SignedString([]byte(JWT_SECRET))
}

func VerifyJWT(jwtToken string) (bool, error) {
	JWT_SECRET := config.GetEnv("JWT_SECRET", "JWT_SECRET")

	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (any, error) {
		return []byte(JWT_SECRET), nil
	})

	if err != nil || !token.Valid {
		helper.LogError(err)
		return false, fmt.Errorf("JWT invalid")
	}

	return true, nil
}
