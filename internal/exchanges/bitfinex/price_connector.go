package bitfinex

import (
	"fmt"

	"github.com/grinply/cryptoapi/trade"
)

type BitfinexPriceConnector struct {
}

func NewPriceConnector() *BitfinexPriceConnector {
	return nil
}

func (conn *BitfinexPriceConnector) GetLatestPrice(tradingPair trade.CurrencyPair) (string, error) {
	return "", fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BitfinexPriceConnector) GetTradingRule(tradingPair trade.CurrencyPair) (*trade.Rule, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BitfinexPriceConnector) GetCandles(tradingPair trade.CurrencyPair, qty int, interval trade.CandleInterval) ([]trade.CandleStick, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *BitfinexPriceConnector) GetPriceFeed(tradingPair trade.CurrencyPair) <-chan trade.CandleStick {
	return nil
}
