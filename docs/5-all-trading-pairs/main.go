package main

import (
	"fmt"

	"github.com/grinply/cryptoapi"
)

func main() {
	infoConnector, err := cryptoapi.GetInfoConnector("binance")

	if err != nil {
		fmt.Printf("Failed to create a connector for the binance exchange with the provided credentials. %v\n", err.Error())
	}

	allTradingPairs, err := infoConnector.GetTradingPairs()

	if err != nil {
		fmt.Printf("Failed to get trading. %v\n", err.Error())
	}

	for _, tradingPair := range allTradingPairs {
		fmt.Printf("%s\n", tradingPair.SimpleName())
	}
}
