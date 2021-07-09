package coinbase

import (
	"fmt"

	"github.com/cryptoapi/trade"
)

type CoinbasePriceConnector struct {
}

func NewPriceConnector() *CoinbasePriceConnector {
	return nil
}

func (conn *CoinbasePriceConnector) GetLatestPrice(tradingPair trade.CurrencyPair) (string, error) {
	return "", fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *CoinbasePriceConnector) GetTradingRule(tradingPair trade.CurrencyPair) (*trade.Rule, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *CoinbasePriceConnector) GetCandles(tradingPair trade.CurrencyPair, qty int, interval trade.CandleInterval) ([]trade.CandleStick, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *CoinbasePriceConnector) GetPriceFeed(tradingPair trade.CurrencyPair) <-chan trade.CandleStick {
	return nil
}
