package order

import (
	"context"

	"github.com/ishchenko-gv/go-example-app/app/user"
)

type Service interface {
	GetOrder(context.Context, OrderID) (*Order, error)
	GetUserOrders(context.Context, user.UserID) ([]Order, error)
	PlaceOrder(context.Context, *Order) error
}

type Repo interface {
	Insert(context.Context, *Order) error
	Find(context.Context, OrderID) (*Order, error)
	FindAllByUserID(context.Context, user.UserID) ([]Order, error)
	Remove(context.Context, OrderID) error
}
