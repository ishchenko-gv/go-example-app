package internal

import (
	"context"

	"github.com/ishchenko-gv/go-example-app/app/order"
	"github.com/ishchenko-gv/go-example-app/app/user"
)

type Service struct {
	Repo order.Repo
}

var _ order.Service = (*Service)(nil)

func (s *Service) GetOrder(ctx context.Context, id order.OrderID) (*order.Order, error) {
	return s.Repo.Find(ctx, id)
}

func (s *Service) GetUserOrders(ctx context.Context, userID user.UserID) ([]order.Order, error) {
	return s.Repo.FindAllByUserID(ctx, userID)
}

func (s *Service) PlaceOrder(ctx context.Context, order *order.Order) error {
	return s.Repo.Insert(ctx, order)
}
