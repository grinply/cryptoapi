package bybit

import (
	"fmt"

	"github.com/grinply/cryptoapi/pkg/trade"
)

type BybitConnector struct{}

func (conn *BybitConnector) GetWalletBalances() ([]trade.Asset, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BybitConnector) GetOrderByID(tradingPair trade.CurrencyPair, orderID string) (*trade.Order, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BybitConnector) OpenOrder(tradingPair trade.CurrencyPair) (string, error) {
	return "", fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BybitConnector) CancelOrder(orderID string) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BybitConnector) CancelAllOrders(tradingPair trade.CurrencyPair) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}
