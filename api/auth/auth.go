package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTCustomClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id int) string {
	claims := &JWTCustomClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}

	return t
}
