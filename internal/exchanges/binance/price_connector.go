package binance

import (
	"fmt"

	"github.com/grinply/cryptoapi/trade"
)

type BinancePriceConnector struct {
}

func NewPriceConnector() *BinancePriceConnector {
	return nil
}

func (conn *BinancePriceConnector) GetLatestPrice(tradingPair trade.CurrencyPair) (string, error) {
	return "", fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BinancePriceConnector) GetTradingRule(tradingPair trade.CurrencyPair) (*trade.Rule, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BinancePriceConnector) GetCandles(tradingPair trade.CurrencyPair, qty int, interval trade.CandleInterval) ([]trade.CandleStick, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BinancePriceConnector) GetPriceFeed(tradingPair trade.CurrencyPair) <-chan trade.CandleStick {
	return nil
}
