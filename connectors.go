package cryptoapi

import (
	"fmt"

	"github.com/cryptoapi/internal/exchanges/binance"
	"github.com/cryptoapi/internal/exchanges/bybit"
	"github.com/cryptoapi/internal/exchanges/coinbase"
	"github.com/cryptoapi/internal/exchanges/ftx"
	"github.com/cryptoapi/internal/exchanges/kucoin"
	"github.com/cryptoapi/internal/exchanges/okex"
	"github.com/cryptoapi/trade"
)

// GetPriceConnector creates a connector that can retrieve pricing information from cryptocurrency exchanges
// a not nil error indicates that the requested exchange is not available
func GetPriceConnector(exchangeName string) (trade.PriceConnector, error) {
	switch exchangeName {
	case "binance":
		return binance.NewPriceConnector(), nil
	case "bybit":
		return bybit.NewPriceConnector(), nil
	case "okex":
		return okex.NewPriceConnector(), nil
	case "ftx":
		return ftx.NewPriceConnector(), nil
	case "coinbase":
		return coinbase.NewPriceConnector(), nil
	case "kucoin":
		return kucoin.NewPriceConnector(), nil
	default:
		return nil, fmt.Errorf("could't find a connector for the provided exchange %s", exchangeName)
	}
}

func GetOrderConnector(exchangeName, apiKey, secretKey string, isTestnet bool) (trade.OrderConnector, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}
