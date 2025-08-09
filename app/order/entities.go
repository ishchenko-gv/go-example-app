package order

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/ishchenko-gv/go-example-app/app/common/money"
	"github.com/ishchenko-gv/go-example-app/app/user"
)

type OrderID uuid.UUID

func (o OrderID) MarshalJSON() ([]byte, error) {
	return json.Marshal(uuid.UUID(o).String())
}

func NewOrderID() OrderID {
	return OrderID(uuid.New())
}

func ZeroOrderID() OrderID {
	return OrderID(uuid.UUID{})
}

func OrderIDFromString(id string) (OrderID, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return ZeroOrderID(), err
	}
	return OrderID(parsed), nil
}

type Order struct {
	ID     OrderID     `json:"id"`
	UserID user.UserID `json:"-"`
	Items  []OrderItem `json:"items"`
}

func NewOrder(userID user.UserID, items []OrderItem) *Order {
	return &Order{
		ID:     NewOrderID(),
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
