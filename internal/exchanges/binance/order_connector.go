package binance

import (
	"context"
	"fmt"
	"strconv"

	"github.com/adshao/go-binance/v2"
	"github.com/grinply/cryptoapi/internal/hash"
	"github.com/grinply/cryptoapi/pkg/trade"
	"github.com/shopspring/decimal"
)

type BinanceConnector struct {
	userID     string //the userID is a hash generate from the apikey + secretkey
	apiKey     string
	secretKey  string
	timeOffset int64
	client     *binance.Client
}

// NewOrderConnector creates a new connector that can open orders on the Binance exchange
func NewOrderConnector(apiKey, secretKey string, isTestnet bool) *BinanceConnector {
	binance.UseTestnet = isTestnet
	var connector = &BinanceConnector{
		apiKey:    apiKey,
		secretKey: secretKey,
		client:    binance.NewClient(apiKey, secretKey),
		userID:    hash.SHA1(apiKey + secretKey)}
	connector.timeOffset, _ = connector.client.NewSetServerTimeService().Do(context.Background())
	return connector
}

func (conn *BinanceConnector) GetWalletBalances() ([]trade.Asset, error) {
	res, err := conn.client.NewGetAccountService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	var balances = make([]trade.Asset, 0, len(res.Balances))

	for _, asset := range res.Balances {
		freeAmount, _ := decimal.NewFromString(asset.Free)
		lockedAmount, _ := decimal.NewFromString(asset.Locked)

		if freeAmount.GreaterThan(decimal.Zero) || lockedAmount.GreaterThan(decimal.Zero) {
			balances = append(balances, trade.Asset{
				Name:      asset.Asset,
				FreeQty:   asset.Free,
				LockedQty: asset.Locked,
			})
		}
	}
	return balances, nil
}

func (conn *BinanceConnector) GetOrderByID(tradingPair trade.CurrencyPair, orderID string) (*trade.Order, error) {
	id, err := strconv.ParseInt(orderID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("the provided id is invalid, couldn't apply the required trandformation from string to integer. %v", err.Error())
	}

	order, err := conn.client.NewGetOrderService().Symbol(tradingPair.SimpleName()).
		OrderID(id).Do(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the requested order with id %v, %v", orderID, err.Error())
	}
	return convertOrder(order, tradingPair), nil
}

func (conn *BinanceConnector) OpenOrder(order trade.Order) (string, error) {
	orderResponse, err := conn.client.NewCreateOrderService().Symbol(order.TradingPair.SimpleName()).
		Side(binance.SideType(order.Direction)).
		Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).
		Quantity(order.BaseQty).
		Price(order.QuotePrice).
		Do(context.Background())

	if err != nil {
		return "", fmt.Errorf("failed to open order on exchange. %v", err.Error())
	}

	return fmt.Sprintf("%v", orderResponse.OrderID), nil
}

func (conn *BinanceConnector) GetAllOpenOrders(tradingPair trade.CurrencyPair) ([]trade.Order, error) {
	openOrders, err := conn.client.NewListOpenOrdersService().Symbol(tradingPair.SimpleName()).
		Do(context.Background())

	if err != nil {
		fmt.Println(err.Error())
		return []trade.Order{}, err
	}

	var orders = make([]trade.Order, 0, len(openOrders))
	for _, binanceOrder := range openOrders {
		orders = append(orders, *convertOrder(binanceOrder, tradingPair))
	}
	return orders, nil
}

func (conn *BinanceConnector) CancelOrder(tradingPair trade.CurrencyPair, orderID string) error {
	_, err := conn.client.NewCancelOrderService().OrderID(111).Symbol(tradingPair.SimpleName()).
		Do(context.Background())
	return err
}

func (conn *BinanceConnector) CancelAllOpenOrders(tradingPair trade.CurrencyPair) error {
	_, err := conn.client.NewCancelOpenOrdersService().Symbol(tradingPair.SimpleName()).Do(context.Background())
	return err
}

func convertOrder(orderBinance *binance.Order, pair trade.CurrencyPair) *trade.Order {
	return &trade.Order{
		ID:          fmt.Sprintf("%v", orderBinance.OrderID),
		TradingPair: pair,
		QuotePrice:  orderBinance.Price,
		BaseQty:     orderBinance.OrigQuantity,
		Direction:   trade.TradeDirection(orderBinance.Side),
		Time:        orderBinance.Time,
	}
}
