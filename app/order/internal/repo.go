package internal

import (
	"context"

	"github.com/ishchenko-gv/go-example-app/app/common/money"
	"github.com/ishchenko-gv/go-example-app/app/order"
	"github.com/ishchenko-gv/go-example-app/app/order/orderid"
	"github.com/ishchenko-gv/go-example-app/app/user"
	"github.com/ishchenko-gv/go-example-app/app/user/userid"
)

type Repo struct{}

type repo interface {
	Insert(context.Context, *order.Order) error
	Find(context.Context, orderid.ID) (*order.Order, error)
	FindAllByUserID(context.Context, userid.ID) ([]order.Order, error)
	Remove(context.Context, orderid.ID) error
}

var _ repo = (*Repo)(nil)

func (r *Repo) Insert(ctx context.Context, order *order.Order) error {
	return nil
}

func (r *Repo) Find(ctx context.Context, id orderid.ID) (*order.Order, error) {
	return order.NewOrder(userid.New(), []order.OrderItem{}), nil
}

func (r *Repo) FindAllByUserID(ctx context.Context, userID userid.ID) ([]order.Order, error) {
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

func (r *Repo) Remove(ctx context.Context, id orderid.ID) error {
	return nil
}
