package trade

import (
	"fmt"
	"strings"
)

const (
	LIMIT  OrderType = "LIMIT"
	MARKET OrderType = "MARKET"
	OCO    OrderType = "OCO"

	BUY  TradeDirection = "BUY"
	SELL TradeDirection = "SELL"
)

// TradeDirection denotes the direction of a trade Order
// example: BUY, SELL
type TradeDirection string

// OrderType is a the type of a order that can be executed in the exchange.
// example: LIMIT, MARKET
type OrderType string

// Order represents a request to buy or sell a amount of cryptocurrency/token in the crypto exchange.
type Order struct {
	ID          string         `json:"id"`
	TradingPair CurrencyPair   `json:"trading_pair"`
	OrderType   OrderType      `json:"order_type"`
	Direction   TradeDirection `json:"direction"`
	BaseQty     string         `json:"base_qty"`
	QuotePrice  string         `json:"quote_price"`
	Time        int64          `json:"time"`
}

// Pair creates a CurrencyPair from a description in the format: BTC/USDT
func Pair(pairDescription string) (CurrencyPair, error) {
	if !strings.Contains(pairDescription, "/") {
		return CurrencyPair{}, fmt.Errorf("the currency pair description must contain a '/' to indicate the separation between base and quote currencies")
	}
	pair := strings.Split(pairDescription, "/")
	return CurrencyPair{
		BaseCurrency:  pair[0],
		QuoteCurrency: pair[1],
	}, nil
}

// CurrencyPair represents a pair of cryptocurrency coins/tokens that can be traded in crypto exchanges.
// example: ETH/BTC, ADA/ETH
type CurrencyPair struct {
	BaseCurrency  string `json:"base_currency"`
	QuoteCurrency string `json:"quote_currency"`
}

// SimpleName returns the currency pair name without separator, for example: ETHBTC, BNBUSDT
func (pair CurrencyPair) SimpleName() string {
	return pair.BaseCurrency + pair.QuoteCurrency
}

func (pair CurrencyPair) Description() string {
	return fmt.Sprintf("%s/%s", pair.BaseCurrency, pair.QuoteCurrency)
}

// Asset represents a cryptocurrency or token that is present in the exchange wallet
type Asset struct {
	Name      string `json:"name"`
	FreeQty   string `json:"free_qty"`
	LockedQty string `json:"locked_qty"`
}
