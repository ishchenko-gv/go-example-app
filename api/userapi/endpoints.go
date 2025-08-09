package userapi

import (
	"encoding/json"
	"net/http"

	"github.com/ishchenko-gv/go-example-app/api/apictx"
	"github.com/ishchenko-gv/go-example-app/app/user"
)

type Endpoints struct {
	UserService user.Service
}

func NewEndpoints(userService user.Service) *Endpoints {
	return &Endpoints{
		UserService: userService,
	}
}

func (ep *Endpoints) Self(w http.ResponseWriter, r *http.Request) (any, error) {
	return apictx.User(r), nil
}

func (ep *Endpoints) PostUser(w http.ResponseWriter, r *http.Request) (any, error) {
	var reqBody UserCreateBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		return nil, err
	}

	u := user.NewUser(reqBody.Email)

	return u, ep.UserService.CreateUser(r.Context(), u, reqBody.Password)
}
