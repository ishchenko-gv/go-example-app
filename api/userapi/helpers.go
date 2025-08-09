package userapi

import (
	"errors"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ishchenko-gv/go-example-app/app/user"
)

func issueJwt(w http.ResponseWriter, usr *user.User) error {
	secret := os.Getenv("SECRET")

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": usr.ID.String(),
		"email":  usr.Email,
	})

	token, err := t.SignedString(secret)
	if err != nil {
		return err
	}

	c := &http.Cookie{
		Name:     "a",
		Value:    token,
		HttpOnly: true,
	}

	http.SetCookie(w, c)

	return nil
}

func VerifyJwt(r *http.Request) (*JwtClaims, error) {
	t, err := r.Cookie("a")
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(t.String(), &JwtClaims{}, func(t *jwt.Token) (any, error) {
		return os.Getenv("SECRET"), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New("can't parse jwt token")
	}

	return claims, nil
}
