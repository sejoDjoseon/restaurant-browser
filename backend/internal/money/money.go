package money

type Currency string

const (
	EUR Currency = "EUR"
)

type Value struct {
	Amount   int      `json:"amount"`
	Currency Currency `json:"currency"`
}
