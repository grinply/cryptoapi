package main

import (
	"fmt"

	"github.com/grinply/cryptoapi"
	"github.com/grinply/cryptoapi/pkg/interval"
	"github.com/grinply/cryptoapi/pkg/trade"
)

func main() {
	priceConnector, err := cryptoapi.GetPriceConnector("binance")

	if err != nil {
		fmt.Printf("Failed to create a connector for the binance exchange with the provided credentials. %v\n", err.Error())
	}

	// Defining the trading pair that will be used
	var tradingPair, _ = trade.Pair("ADA/ETH")
	var candlestickQty = 350

	candles, err := priceConnector.GetCandles(tradingPair, candlestickQty, interval.ThreeMinutes)

	if err != nil {
		fmt.Printf("Failed to retrieve the desired candles for the trading pair %s, %v\n", tradingPair.SimpleName(), err.Error())
		return
	}

	//Print all retrieved candlesticks
	for _, candle := range candles {
		fmt.Printf("%v\n", candle)
	}
}
