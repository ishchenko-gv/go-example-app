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
		h.Middleware.PanicRecoveryMiddleware,
		h.Middleware.LoggingMiddleware,
		h.Middleware.JsonContentMiddleware,
		h.Middleware.AuthMiddleware,
	)(mux)

	return wrappedMux
}

func (h *handler) setupUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users/self", ep(h.User.GetSelf))
	mux.HandleFunc("POST /users/register", ep(h.User.PostRegister))
	mux.HandleFunc("POST /users/login", ep(h.User.PostLogin))
}

func (h *handler) setupOrderRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /orders", ep(h.Order.GetOrders))
	mux.HandleFunc("GET /orders/{id}", ep(h.Order.GetOrder))
	mux.HandleFunc("POST /orders", ep(h.Order.PostOrder))
}
