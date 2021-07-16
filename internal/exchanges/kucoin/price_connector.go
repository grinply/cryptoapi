package kucoin

import (
	"fmt"

	"github.com/grinply/cryptoapi/pkg/trade"
)

type KuCoinPriceConnector struct {
}

func NewPriceConnector() *KuCoinPriceConnector {
	return nil
}

func (conn *KuCoinPriceConnector) GetLatestPrice(tradingPair trade.CurrencyPair) (string, error) {
	return "", fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *KuCoinPriceConnector) GetTradingRule(tradingPair trade.CurrencyPair) (*trade.Rule, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *KuCoinPriceConnector) GetCandles(tradingPair trade.CurrencyPair, qty int, interval trade.CandleTime) ([]trade.CandleStick, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *KuCoinPriceConnector) GetPriceFeed(tradingPair trade.CurrencyPair) <-chan trade.CandleStick {
	return nil
}
