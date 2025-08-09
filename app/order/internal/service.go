package internal

import (
	"context"

	"github.com/ishchenko-gv/go-example-app/app/order"
	"github.com/ishchenko-gv/go-example-app/app/order/orderid"
	"github.com/ishchenko-gv/go-example-app/app/user/userid"
)

type Service struct {
	Repo repo
}

var _ order.Service = (*Service)(nil)

func (s *Service) GetOrder(ctx context.Context, id orderid.ID) (*order.Order, error) {
	return s.Repo.Find(ctx, id)
}

func (s *Service) GetUserOrders(ctx context.Context, userID userid.ID) ([]order.Order, error) {
	return s.Repo.FindAllByUserID(ctx, userID)
}

func (s *Service) PlaceOrder(ctx context.Context, order *order.Order) error {
	return s.Repo.Insert(ctx, order)
}
