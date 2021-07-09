package okex

import (
	"fmt"

	"github.com/cryptoapi/trade"
)

type OkexPriceConnector struct {
}

func NewPriceConnector() *OkexPriceConnector {
	return nil
}

func (conn *OkexPriceConnector) GetLatestPrice(tradingPair trade.CurrencyPair) (string, error) {
	return "", fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *OkexPriceConnector) GetTradingRule(tradingPair trade.CurrencyPair) (*trade.Rule, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *OkexPriceConnector) GetCandles(tradingPair trade.CurrencyPair, qty int, interval trade.CandleInterval) ([]trade.CandleStick, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *OkexPriceConnector) GetPriceFeed(tradingPair trade.CurrencyPair) <-chan trade.CandleStick {
	return nil
}
