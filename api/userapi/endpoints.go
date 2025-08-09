package userapi

import (
	"net/http"

	"github.com/ishchenko-gv/go-example-app/api/apictx"
)

type Endpoints struct{}

func NewEndpoints() *Endpoints {
	return &Endpoints{}
}

func (ep *Endpoints) Self(w http.ResponseWriter, r *http.Request) (any, error) {
	return apictx.User(r), nil
}
