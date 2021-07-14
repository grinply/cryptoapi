package cryptoapi

import (
	"fmt"
	"strings"

	"github.com/grinply/cryptoapi/internal/exchanges/binance"
	"github.com/grinply/cryptoapi/trade"
)

// GetPriceConnector creates a connector that can retrieve pricing information for the specified cryptocurrency exchange
// a not nil error indicates that a connector for the exchange is not available
func GetPriceConnector(exchangeName string) (trade.PriceConnector, error) {
	switch strings.ToLower(exchangeName) {
	case "binance":
		return binance.NewPriceConnector(), nil
	default:
		return nil, fmt.Errorf("could't find a price connector for the provided exchange %s", exchangeName)
	}
}

// GetOrderConnector creates a connector that can execute orders and access private information for the specified cryptocurrency exchange
// a not nil error indicates wrong authentication information or that a connector for the exchange is not available
func GetOrderConnector(exchangeName, apiKey, secretKey string, isTestnet bool) (trade.OrderConnector, error) {
	switch strings.ToLower(exchangeName) {
	case "binance":
		return binance.NewOrderConnector(apiKey, secretKey, isTestnet), nil
	default:
		return nil, fmt.Errorf("could't find a order connector for the provided exchange %s", exchangeName)
	}
}
