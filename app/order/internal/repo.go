package internal

import (
	"context"

	"github.com/ishchenko-gv/go-example-app/app/common/money"
	"github.com/ishchenko-gv/go-example-app/app/order"
	"github.com/ishchenko-gv/go-example-app/app/user"
)

type Repo struct{}

var _ order.Repo = (*Repo)(nil)

func (r *Repo) Insert(ctx context.Context, order *order.Order) error {
	return nil
}

func (r *Repo) Find(ctx context.Context, id order.OrderID) (*order.Order, error) {
	return order.NewOrder(user.NewUserID(), []order.OrderItem{}), nil
}

func (r *Repo) FindAllByUserID(ctx context.Context, userID user.UserID) ([]order.Order, error) {
	orders := []order.Order{
		*order.NewOrder(user.NewUser().ID, []order.OrderItem{
			*order.NewOrderItem("Item 1", *money.NewMoney(100, money.Euro)),
			*order.NewOrderItem("Item 2", *money.NewMoney(150, money.Euro)),
			*order.NewOrderItem("Item 3", *money.NewMoney(200, money.Euro)),
		}),
		*order.NewOrder(user.NewUser().ID, []order.OrderItem{
			*order.NewOrderItem("Item 1", *money.NewMoney(300, money.Euro)),
			*order.NewOrderItem("Item 2", *money.NewMoney(350, money.Euro)),
			*order.NewOrderItem("Item 3", *money.NewMoney(400, money.Euro)),
		}),
	}

	return orders, nil
}

func (r *Repo) Remove(ctx context.Context, id order.OrderID) error {
	return nil
}
