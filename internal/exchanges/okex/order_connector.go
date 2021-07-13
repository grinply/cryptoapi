package okex

import (
	"fmt"

	"github.com/grinply/cryptoapi/trade"
)

type OkexConnector struct{}

func (conn *OkexConnector) GetWalletBalances() ([]trade.Asset, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *OkexConnector) GetOrderByID(tradingPair trade.CurrencyPair, orderID string) (*trade.Order, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *OkexConnector) OpenOrder(tradingPair trade.CurrencyPair) (string, error) {
	return "", fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *OkexConnector) CancelOrder(orderID string) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *OkexConnector) CancelAllOrders(tradingPair trade.CurrencyPair) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}
