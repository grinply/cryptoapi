[![Go Report Card](https://goreportcard.com/badge/github.com/grinply/universal-crypto-api?style=flat-square)](https://goreportcard.com/report/github.com/grinply/universal-crypto-api)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE.md)
[![Go Reference](https://pkg.go.dev/badge/github.com/grinply/cryptoapi.svg)](https://pkg.go.dev/github.com/victorl2/kate-backtester)
# Universal Crypto API

**Universal API** interface in Golang that aims to abstract the connection with the [**biggest cryptocurrency exchanges by Volume**](https://coinmarketcap.com/rankings/exchanges/). The API supports both **spot** and **derivative** markets. 

## Trading Pairs in Cryptocurrency
Understanding how trading pairs work is important for effect use of the universal crypto api, but also trading in general.
In cryptocurrency, _“trading pairs” or “cryptocurrency pairs”_ are assets that can be traded for each other on an exchange — for example Bitcoin/Litecoin **(BTC/LTC)** and Ethereum/Bitcoin Cash **(ETH/BCH)**. Trading pairs lets you compare costs between different cryptocurrencies.


### Operating in an trading pair
Currency pairs compare the value of one currency to another—the **base currency** _(or the first one)_ versus the second or the **quote currency**. It indicates how much of the **quote currency** is needed to purchase **one unit** of the **base currency**.


For example if we are trading the **ETH/BTC** pair and see the following chart:

![ETH/BTC price chart](https://pbs.twimg.com/media/E4_dvSgXMAILbVY.jpg)

The latest price denoted in the chart is **0.060943**, _what does this mean ?_ this can be translate as **0.060943 BTC** _(quote currency)_ can BUY **1 ETH** _(base currency)_, in all operations you either buy or sell the base currency _(ETH in this case)_, the quote currency _(BTC in the example)_ is only used to price the operations.

A example of some operations can be:

+ **Initial balance:** 1 BTC and 2 ETH
+ **BUY**  2.0 ETH for 0.121886 BTC _( price of 0.060943)_
+ **updated balance:** 4 ETH and 0.878114 BTC
+ **SELL** 0.5 ETH for 0.035 BTC _( price of 0.07000)_
+ **updated balance:** 4 ETH and 0.878114 BTC
+ **BUY**  0.5 ETH for 0.31 BTC _( price of 0.06)_
+ **final balance:** 4 ETH and 0.878114 BTC

