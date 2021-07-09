package binance

import (
	"fmt"

	"github.com/cryptoapi/trade"
)

type BinanceConnector struct{}

func (conn *BinanceConnector) GetWalletBalances() ([]trade.Asset, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BinanceConnector) GetOrderByID(tradingPair trade.CurrencyPair, orderID string) (*trade.Order, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BinanceConnector) OpenOrder(tradingPair trade.CurrencyPair) (string, error) {
	return "", fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BinanceConnector) CancelOrder(orderID string) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BinanceConnector) CancelAllOrders(tradingPair trade.CurrencyPair) error {
	return fmt.Errorf("NOT IMPLEMENTED")
}
