package bybit

import (
	"fmt"

	"github.com/cryptoapi/trade"
)

type BybitPriceConnector struct {
}

func NewPriceConnector() *BybitPriceConnector {
	return nil
}

func (conn *BybitPriceConnector) GetLatestPrice(tradingPair trade.CurrencyPair) (string, error) {
	return "", fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BybitPriceConnector) GetTradingRule(tradingPair trade.CurrencyPair) (*trade.Rule, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BybitPriceConnector) GetCandles(tradingPair trade.CurrencyPair, qty int, interval trade.CandleInterval) ([]trade.CandleStick, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BybitPriceConnector) GetPriceFeed(tradingPair trade.CurrencyPair) <-chan trade.CandleStick {
	return nil
}
