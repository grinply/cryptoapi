package bitfinex

import (
	"fmt"

	"github.com/cryptoapi/trade"
)

type BitfinexConnector struct{}

func (conn *BitfinexConnector) GetWalletBalances() ([]trade.Asset, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BitfinexConnector) GetOrderByID(tradingPair trade.CurrencyPair, orderID string) (*trade.Order, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BitfinexConnector) OpenOrder(tradingPair trade.CurrencyPair) (string, error) {
	return "", fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BitfinexConnector) CancelOrder(orderID string) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BitfinexConnector) CancelAllOrders(tradingPair trade.CurrencyPair) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}
