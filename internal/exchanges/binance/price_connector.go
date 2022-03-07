package binance

import (
	"context"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/grinply/cryptoapi/pkg/interval"
	"github.com/grinply/cryptoapi/pkg/trade"
)

const (
	candles_per_request int64 = 500
	minute_in_ms        int64 = 60 * 1e3
)

type BinancePriceConnector struct {
	timeOffset int64
	client     *binance.Client
}

func NewPriceConnector() *BinancePriceConnector {
		var connector = &BinancePriceConnector{
			client:    binance.NewClient("", "")
		}
	}
	connector.timeOffset, _ = connector.client.NewSetServerTimeService().Do(context.Background())
	return connector
}

func (conn *BinancePriceConnector) GetLatestPrice(tradingPair trade.CurrencyPair) (string, error) {
	stats, err := conn.client.NewListPriceChangeStatsService().Symbol(tradingPair.SimpleName()).Do(context.Background())

	if err != nil {
		return "", err
	}

	return stats[-1].LastPrice //Validar se é o 0 ou o -1 o último status.
}


func (conn *BinancePriceConnector) GetPriceFeed(tradingPair trade.CurrencyPair) <-chan trade.CandleStick {
	return nil
}

func (conn *BinancePriceConnector) GetCandles(tradingPair trade.CurrencyPair, qty int, interval trade.CandleTime) ([]trade.CandleStick, error) {
	var currentTime = time.Now().Unix()*1e3 - conn.timeOffset
	var priceCandles = make([]trade.CandleStick, 0, qty)

	for i := qty; i > 0; i -= int(candles_per_request) {
		var startCandles = currentTime - int64(i)*interval.Minutes()*minute_in_ms
		var endCandles = startCandles + interval.Minutes()*minute_in_ms*candles_per_request
		var candles, _ = conn.GetCandlesInRange(tradingPair, interval, startCandles, endCandles)
		priceCandles = append(priceCandles, candles...)
	}

	return priceCandles, nil
}

func (conn *BinancePriceConnector) GetCandlesInRange(pair trade.CurrencyPair, interval trade.CandleTime, startTime, endTime int64) ([]trade.CandleStick, error) {
	klines, err := conn.client.NewKlinesService().Symbol(pair.SimpleName()).
		Interval(interval.Description()).StartTime(startTime).EndTime(endTime).Do(context.Background())

	var candles = make([]trade.CandleStick, 0, len(klines))

	if err != nil {
		return []trade.CandleStick{}, err
	}

	for i, kline := range klines {
		candles = append(candles, trade.CandleStick{
			TradingPair: pair,
			StartTime:   strconv.Itoa(int(startTime) + i*int(interval.Minutes()*minute_in_ms)),
			Open:        kline.Open,
			High:        kline.High,
			Low:         kline.Low,
			Close:       kline.Close,
			Volume:      kline.Volume,
			IsClosed:    true,
		})
	}

	return candles, nil
}
