package order

import (
	"context"

	"github.com/ishchenko-gv/go-example-app/app/order/orderid"
	"github.com/ishchenko-gv/go-example-app/app/user/userid"
)

type Service interface {
	GetOrder(context.Context, orderid.ID) (*Order, error)
	GetUserOrders(context.Context, userid.ID) ([]Order, error)
	PlaceOrder(context.Context, *Order) error
}
