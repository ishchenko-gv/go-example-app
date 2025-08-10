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

func (ep *Endpoints) PostRegister(w http.ResponseWriter, r *http.Request) (any, error) {
	var reqBody UserCreateBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		return nil, err
	}

	usr := user.NewUser(reqBody.Email)

	err = issueJwt(w, usr)
	if err != nil {
		return nil, err
	}

	return usr, ep.UserService.CreateUser(r.Context(), usr, reqBody.Password)
}

func (ep *Endpoints) PostLogin(w http.ResponseWriter, r *http.Request) (any, error) {
	var reqBody UserCreateBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		return nil, err
	}

	usr, err := ep.UserService.AuthenticateByEmail(r.Context(), reqBody.Email, reqBody.Password)
	if err != nil {
		return nil, err
	}

	err = issueJwt(w, usr)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
