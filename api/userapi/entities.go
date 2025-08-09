package userapi

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/ishchenko-gv/go-example-app/app/user/userid"
)

type UserCreateBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtClaims struct {
	jwt.RegisteredClaims
	ID    userid.ID `json:"id"`
	Email string    `json:"email"`
}
