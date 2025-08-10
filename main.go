package main

import (
	"fmt"

	"github.com/ishchenko-gv/go-example-app/api"
	"github.com/ishchenko-gv/go-example-app/app/order/orderfactory"
	"github.com/ishchenko-gv/go-example-app/app/user/userfactory"
	"github.com/ishchenko-gv/go-example-app/db"
	"github.com/ishchenko-gv/go-example-app/env"
)

func main() {
	env.Setup()

	db.Connect()
	defer db.DB.Close()

	userRepo := userfactory.NewRepo(db.DB)
	userService := userfactory.NewService(userRepo)

	orderRepo := orderfactory.NewRepo(db.DB)
	orderService := orderfactory.NewService(orderRepo)

	middlerware := api.NewMiddleware(userService)

	handler := api.NewHandler(
		middlerware,
		userService,
		orderService,
	).Setup()

	server := api.NewServer(handler)

	fmt.Println("Starting http server...")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
