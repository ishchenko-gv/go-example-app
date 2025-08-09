package api

import (
	"net/http"

	"github.com/ishchenko-gv/go-example-app/api/orderapi"
	"github.com/ishchenko-gv/go-example-app/api/userapi"
	"github.com/ishchenko-gv/go-example-app/app/order"
	"github.com/ishchenko-gv/go-example-app/app/user"
)

type handler struct {
	User  *userapi.Endpoints
	Order *orderapi.Endpoints
}

type Endpoint func(http.ResponseWriter, *http.Request) (any, error)

func NewHandler(
	userService user.Service,
	orderService order.Service,
) *handler {
	return &handler{
		User:  userapi.NewEndpoints(userService),
		Order: orderapi.NewEndpoints(orderService),
	}
}

func (h *handler) Setup() *http.ServeMux {
	mux := http.NewServeMux()

	h.setupUserRoutes(mux)
	h.setupOrderRoutes(mux)

	return mux
}

func (h *handler) setupUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users/self", wrap(h.User.Self))
	mux.HandleFunc("POST /users", wrap(h.User.PostUser))
}

func (h *handler) setupOrderRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /orders", wrap(h.Order.GetOrders))
	mux.HandleFunc("GET /orders/{id}", wrap(h.Order.GetOrder))
	mux.HandleFunc("POST /orders", wrap(h.Order.PostOrder))
}
