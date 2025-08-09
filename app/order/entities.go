package order

import (
	"github.com/google/uuid"
	"github.com/ishchenko-gv/go-example-app/app/common/money"
	"github.com/ishchenko-gv/go-example-app/app/order/orderid"
	"github.com/ishchenko-gv/go-example-app/app/user/userid"
)

type Order struct {
	ID     orderid.ID  `json:"id"`
	UserID userid.ID   `json:"-"`
	Items  []OrderItem `json:"items"`
}

func NewOrder(userID userid.ID, items []OrderItem) *Order {
	return &Order{
		ID:     orderid.New(),
		UserID: userID,
		Items:  items,
	}
}

type OrderItemID uuid.UUID

func NewOrderItemID() OrderItemID {
	return OrderItemID(uuid.New())
}

type OrderItem struct {
	ID    uuid.UUID   `json:"id"`
	Title string      `json:"title"`
	Price money.Money `json:"price"`
}

func NewOrderItem(title string, price money.Money) *OrderItem {
	return &OrderItem{
		ID:    uuid.UUID(NewOrderItemID()),
		Title: title,
		Price: price,
	}
}
