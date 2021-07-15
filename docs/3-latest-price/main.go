package main

import (
	"fmt"

	"github.com/grinply/cryptoapi"
	"github.com/grinply/cryptoapi/pkg/trade"
)

func main() {
	priceConnector, err := cryptoapi.GetPriceConnector("binance")

	if err != nil {
		fmt.Printf("Failed to create a connector for the binance exchange with the provided credentials. %v\n", err.Error())
	}

	// Defining the trading pair that will be used
	var tradingPair, _ = trade.Pair("ADA/ETH")

	price, err := priceConnector.GetLatestPrice(tradingPair)

	if err != nil {
		fmt.Printf("Failed to retrieve the latest price for the trading pair %s, %v\n", tradingPair.SimpleName(), err.Error())
		return
	}

	fmt.Printf("The latest price in the pair %v is 1%v for %v%v\n", tradingPair.Description(), tradingPair.BaseCurrency,
		price, tradingPair.QuoteCurrency)
}
