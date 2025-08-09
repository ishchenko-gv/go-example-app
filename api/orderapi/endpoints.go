package orderapi

import (
	"net/http"

	"github.com/ishchenko-gv/go-example-app/api/apictx"
	"github.com/ishchenko-gv/go-example-app/app/order"
	"github.com/ishchenko-gv/go-example-app/app/order/orderid"
)

type Endpoints struct {
	OrderService order.Service
}

func NewEndpoints(orderService order.Service) *Endpoints {
	return &Endpoints{
		OrderService: orderService,
	}
}

func (ep *Endpoints) GetOrders(w http.ResponseWriter, r *http.Request) (any, error) {
	user := apictx.User(r)

	return ep.OrderService.GetUserOrders(r.Context(), user.ID)
}

func (ep *Endpoints) GetOrder(w http.ResponseWriter, r *http.Request) (any, error) {
	id, err := orderid.FromString(r.PathValue("id"))
	if err != nil {
		return nil, err
	}

	return ep.OrderService.GetOrder(r.Context(), id)
}

func (ep *Endpoints) PostOrder(w http.ResponseWriter, r *http.Request) (any, error) {
	user := apictx.User(r)
	ord := order.NewOrder(user.ID, []order.OrderItem{})

	return ord, ep.OrderService.PlaceOrder(r.Context(), ord)
}
