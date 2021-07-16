package ftx

import (
	"fmt"

	"github.com/grinply/cryptoapi/pkg/trade"
)

type FTXPriceConnector struct {
}

func NewPriceConnector() *FTXPriceConnector {
	return nil
}

func (conn *FTXPriceConnector) GetLatestPrice(tradingPair trade.CurrencyPair) (string, error) {
	return "", fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *FTXPriceConnector) GetTradingRule(tradingPair trade.CurrencyPair) (*trade.Rule, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *FTXPriceConnector) GetCandles(tradingPair trade.CurrencyPair, qty int, interval trade.CandleTime) ([]trade.CandleStick, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (conn *FTXPriceConnector) GetPriceFeed(tradingPair trade.CurrencyPair) <-chan trade.CandleStick {
	return nil
}
