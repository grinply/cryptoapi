package binance

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/grinply/cryptoapi/pkg/trade"
)

type BinanceInfoConnector struct {
	client *binance.Client
}

var exchangeRules map[string]trade.Rule
var tradingPairs []trade.CurrencyPair

func NewBinanceInfoConnector() *BinanceInfoConnector {
	return &BinanceInfoConnector{
		client: binance.NewClient("", ""),
	}
}

// GetTradingPairs returns all avaliable pairs on the exchange
func (conn *BinanceInfoConnector) GetTradingPairs() ([]trade.CurrencyPair, error) {
	if tradingPairs != nil {
		return tradingPairs, nil
	}

	if err := conn.loadExchangeRules(); err != nil {
		return nil, fmt.Errorf("failed to load trading pairs from Binance, %v", err.Error())
	}

	tradingPairs = make([]trade.CurrencyPair, 10)
	for _, rule := range exchangeRules {
		tradingPairs = append(tradingPairs, rule.TradingPair)
	}

	return tradingPairs, nil
}

func (conn *BinanceInfoConnector) GetTradingRule(tradingPair trade.CurrencyPair) (*trade.Rule, error) {
	if rule, ok := exchangeRules[tradingPair.SimpleName()]; ok {
		return &rule, nil
	} else if len(exchangeRules) > 0 {
		return nil, fmt.Errorf("the exchange didn't provide a rule for the desired trading pair %v, check if the name is correct",
			tradingPair.Description())
	}
	conn.loadExchangeRules()
	return conn.GetTradingRule(tradingPair)
}

func (conn *BinanceInfoConnector) loadExchangeRules() error {
	exchangeInfo, err := conn.client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		return fmt.Errorf("failed to retrieve exchange information with trade rules, %v", err.Error())
	}

	exchangeRules = make(map[string]trade.Rule)

	for _, exchangeRule := range exchangeInfo.Symbols {
		pair, _ := trade.Pair(exchangeRule.BaseAsset + "/" + exchangeRule.QuoteAsset)
		newRule := trade.Rule{
			TradingPair:         pair,
			BaseAssetPrecision:  exchangeRule.BaseAssetPrecision,
			QuoteAssetPrecision: exchangeRule.QuotePrecision,
			MaxOrderQty:         200,
			MinPriceMovement:    exchangeRule.PriceFilter().TickSize,
			BaseStepSize:        exchangeRule.LotSizeFilter().StepSize,
			MinLotSize:          exchangeRule.LotSizeFilter().MinQuantity,
			MinNotionalValue:    exchangeRule.MinNotionalFilter().MinNotional,
		}
		exchangeRules[pair.SimpleName()] = newRule
	}

	return nil
}
