[![Go Report Card](https://goreportcard.com/badge/github.com/grinply/universal-crypto-api?style=flat-square)](https://goreportcard.com/report/github.com/grinply/universal-crypto-api)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE.md)
[![Go Reference](https://pkg.go.dev/badge/github.com/grinply/cryptoapi.svg)](https://pkg.go.dev/github.com/victorl2/kate-backtester)
# CryptoAPI

**CryptoAPI** is a interface in Golang that aims to abstract the connection with the [**biggest cryptocurrency exchanges by Volume**](https://coinmarketcap.com/rankings/exchanges/). The API supports both **spot** and **derivative** markets. 



# Usage

First you need to install the **CryptoAPI** in your golang project, run `go get github.com/grinply/cryptoapi` to add the dependency.

The CryptoAPI provides **2 interfaces** that abstracts access to exchanges. [**OrderConnector**]((trade/order_connector.go)) allows users to execute [**orders**](https://www.tradingpedia.com/bitcoin-guide/what-types-of-orders-to-trade-bitcoin-on-crypto-exchanges-are-there/) and access private information _(such as asset balances)_:

```go
package trade

type OrderConnector interface {
	GetWalletBalances() ([]Asset, error)

	GetOrderByID(tradingPair CurrencyPair, orderID string) (*Order, error)

	GetAllOpenOrders(tradingPair CurrencyPair) ([]Order, error)

	OpenOrder(orderToOpen Order) (string, error)

	CancelOrder(tradingPair CurrencyPair, orderID string) error

	CancelAllOpenOrders(tradingPair CurrencyPair) error
}

```

The [**PriceConnector**](trade/price_connector.go) provides access to **price data** without the need for authentication:

```go

type PriceConnector interface {
	GetLatestPrice(tradingPair CurrencyPair) (string, error)

	GetTradingRule(tradingPair CurrencyPair) (*Rule, error)

	GetCandles(tradingPair CurrencyPair, qty int, interval CandleInterval) 
        ([]CandleStick, error)

	GetPriceFeed(tradingPair CurrencyPair) <-chan CandleStick
}
```

To make use of both connectors you just need to provide the required information to a function, a implementation for the requested exchange will be provided for your use:

```go
package main

import "github.com/cryptoapi"

func main(){
    var exchangeName = "binance"
    priceConnector, err := cryptoapi.GetPriceConnector(exchangeName)

    if err != nil {
        fmt.Println("could not generate a priceConnector")
        return
    }  

    var apiKey = "my_api_key"
    var secretKey = "my_secret_key"
    var isTestNet = false

    orderConnector, err := cryptoapi.GetOrderConnector(exchangeName, apiKey, secretKey, isTestNet) 
    if err != nil {
        fmt.Println("could not generate a orderConnector")
        return
    }
}
```

[_Check here a list of examples for general CryptoAPI usage](docs/)

