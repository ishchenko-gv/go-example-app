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
		"id":    usr.ID.String(),
		"email": usr.Email,
	})

	token, err := t.SignedString([]byte(secret))
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

func VerifyJwt(t string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(t, &JwtClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("SECRET")), nil
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
