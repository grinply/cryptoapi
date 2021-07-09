package trade

const (
	LIMIT  OrderType = "LIMIT"
	MARKET OrderType = "MARKET"
)

// OrderType is a the type of a order that can be executed in the exchange.
// example: LIMIT, MARKET
type OrderType string

// Order represents a request to buy or sell a amount of cryptocurrency/token in the crypto exchange.
type Order struct {
	OrderType OrderType `json:"order_type"`
}

// CurrencyPair represents a pair of cryptocurrency coins/tokens that can be traded in crypto exchanges.
// example: ETH/BTC, ADA/ETH
type CurrencyPair struct {
	BaseCurrency  string `json:"base_currency"`
	QuoteCurrency string `json:"quote_currency"`
}

// Asset represents a cryptocurrency or token that is present in the exchange wallet
type Asset struct {
	Name         string `json:"name"`
	Amount       string `json:"amount"`
	LockedAmount string `json:"locked_amount"`
}
