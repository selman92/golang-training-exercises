package money

func New(amount int, currency string) Money {
	return Money{
		Amount:   amount,
		Currency: currency,
	}
}

type Money struct {
	Amount   int
	Currency string
}
