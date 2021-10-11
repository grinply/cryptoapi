[![Go Report Card](https://goreportcard.com/badge/github.com/grinply/universal-crypto-api?style=flat-square)](https://goreportcard.com/report/github.com/grinply/universal-crypto-api)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE.md)
[![Go Reference](https://pkg.go.dev/badge/github.com/grinply/cryptoapi.svg)](https://pkg.go.dev/github.com/victorl2/kate-backtester)
# CryptoAPI

**CryptoAPI** is a interface in Golang that aims to abstract the connection with [cryptocurrency exchanges](https://coinmarketcap.com/rankings/exchanges/). The API supports both **spot** and **derivative** markets. 



# Usage

First you need to install the **CryptoAPI** in your project, run `go get github.com/grinply/cryptoapi` to add the dependency.

Three main interfaces are provided with abstractions to access exchanges in a unified way. [**OrderConnector**]((trade/order_connector.go)) allows users to execute [**orders**](https://www.tradingpedia.com/bitcoin-guide/what-types-of-orders-to-trade-bitcoin-on-crypto-exchanges-are-there/) and access private information _(such as asset balances)_:

```go
package trade

type OrderConnector interface {
	FindOrderByID(tradingPair CurrencyPair, orderID string) (Order, error)

	OpenOrders(tradingPair CurrencyPair) ([]Order, error)

	NewOpenOrder(orderToOpen Order) (string, error)

	CancelOrder(tradingPair CurrencyPair, orderID string) error

	CancelOpenOrders(tradingPair CurrencyPair) error
}

```

The [**PriceConnector**](trade/price_connector.go) provides access to **price data** without the need for authentication:

```go

type PriceConnector interface {
	LatestPrice(tradingPair CurrencyPair) (string, error)

	Candles(tradingPair CurrencyPair, qty int, interval CandleInterval) 
        ([]CandleStick, error)

	PriceFeed(tradingPair CurrencyPair) <-chan CandleStick
}
```

The [**InfoConnector**](trade/info_connector.go) provides access to **general exchange information** about what is avaliable

```go
type InfoConnector interface {
	TradingPairs() ([]CurrencyPair, error)

	TradingRules(tradingPair CurrencyPair) (Rule, error)

	CoinsBalance() ([]Asset, error)
}
```

To make use of both connectors you just need to provide the required information to a function, a implementation for the requested exchange will be provided for your use:

```go
package main

import (
	"fmt"
	"github.com/grinply/cryptoapi"
)

func main() {
    //replace the keys  with your own values.
    var apiKey = "my_api_key"
    var secretKey = "my_secret_key"
    var exchange = "binance"

	connector, err := cryptoapi.OrderConnector(exchange, apiKey, secretKey, false)

	if err != nil {
		fmt.Printf("Fail to create a connector with the provided credentials. %v\n", err.Error())
		return
	}

	//Print the amount of each asset present in the exchange wallet
	if assetsBalance, err := connector.WalletBalances(); err == nil {
		for _, asset := range assetsBalance {
			fmt.Printf("%s - available: %s | locked: %s\n", 
				asset.Name, asset.FreeQty, asset.LockedQty)
		}
	}
}
```

[_Check here a list of **examples** for general CryptoAPI usage_](docs/)
