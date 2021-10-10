package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/grinply/cryptoapi"
)

// In this example we will load the apikey and secret from a json connect to Binance and check the wallet balance for each asset
// NOTE: to reproduce this example you must replace the keys in the json file with your own values.
func main() {
	config, readErr := ioutil.ReadFile("config.json")
	if readErr != nil {
		fmt.Printf("Failed to read json content from config.json. %v", readErr.Error())
		return
	}

	var data map[string]string
	if err := json.Unmarshal(config, &data); err != nil {
		fmt.Println("Failed to parse authentication information from config.json")
		return
	}

	connector, err := cryptoapi.GetOrderConnector(data["exchange"], data["api_key"], data["secret_key"], false)

	if err != nil {
		fmt.Printf("Failed to create a connector for the binance exchange with the provided credentials. %v\n", err.Error())
		return
	}

	//Print the amount of each asset present in the exchange wallet
	if assetsBalance, err := connector.GetWalletBalances(); err == nil {
		for _, asset := range assetsBalance {
			fmt.Printf("%s - available: %s | locked: %s\n", asset.Name, asset.FreeQty, asset.LockedQty)
		}
	}
}
