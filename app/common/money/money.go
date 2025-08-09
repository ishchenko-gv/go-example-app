package money

type Money struct {
	Amount   int      `json:"amount"`
	Currency Currency `json:"ccy"`
}

type Currency string

var Euro Currency = "EUR"

func NewMoney(amount int, ccy Currency) *Money {
	return &Money{
		Amount:   amount,
		Currency: ccy,
	}
}
