package api

import (
	"net/http"

	"github.com/ishchenko-gv/go-example-app/api/orderapi"
	"github.com/ishchenko-gv/go-example-app/api/userapi"
	"github.com/ishchenko-gv/go-example-app/app/order"
	"github.com/ishchenko-gv/go-example-app/app/user"
)

type handler struct {
	Middleware *middleware
	User       *userapi.Endpoints
	Order      *orderapi.Endpoints
}

type Endpoint func(http.ResponseWriter, *http.Request) (any, error)

func NewHandler(
	middleware *middleware,
	userService user.Service,
	orderService order.Service,
) *handler {
	return &handler{
		Middleware: middleware,
		User:       userapi.NewEndpoints(userService),
		Order:      orderapi.NewEndpoints(orderService),
	}
}

func (h *handler) Setup() http.Handler {
	mux := http.NewServeMux()

	h.setupUserRoutes(mux)
	h.setupOrderRoutes(mux)

	wrappedMux := chainMiddlewares(
		h.Middleware.JsonContentMiddleware,
		h.Middleware.AuthMiddleware,
		h.Middleware.LoggingMiddleware,
	)(mux)

	return wrappedMux
}

func (h *handler) setupUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users/self", wrap(h.User.Self))
	mux.HandleFunc("POST /users/register", wrap(h.User.PostRegister))
	mux.HandleFunc("POST /users/login", wrap(h.User.PostLogin))
}

func (h *handler) setupOrderRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /orders", wrap(h.Order.GetOrders))
	mux.HandleFunc("GET /orders/{id}", wrap(h.Order.GetOrder))
	mux.HandleFunc("POST /orders", wrap(h.Order.PostOrder))
}
