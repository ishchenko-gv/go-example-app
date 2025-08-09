package api

import (
	"net/http"
)

func NewServer(handler http.Handler) http.Server {
	return http.Server{
		Addr:    ":3000",
		Handler: handler,
	}
}
