package coinbase

import (
	"fmt"

	"github.com/cryptoapi/trade"
)

type CoinbaseConnector struct{}

func (conn *CoinbaseConnector) GetWalletBalances() ([]trade.Asset, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *CoinbaseConnector) GetOrderByID(tradingPair trade.CurrencyPair, orderID string) (*trade.Order, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *CoinbaseConnector) OpenOrder(tradingPair trade.CurrencyPair) (string, error) {
	return "", fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *CoinbaseConnector) CancelOrder(orderID string) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *CoinbaseConnector) CancelAllOrders(tradingPair trade.CurrencyPair) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}
