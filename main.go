package main

import (
	"fmt"

	"github.com/ishchenko-gv/go-example-app/api"
	"github.com/ishchenko-gv/go-example-app/app/order/orderfactory"
)

func main() {
	orderRepo := orderfactory.NewRepo()
	orderService := orderfactory.NewService(orderRepo)

	handler := api.NewHandler(orderService).Setup()
	server := api.NewServer(handler)

	fmt.Println("Staring http server...")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
