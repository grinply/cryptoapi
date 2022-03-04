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

type BinanceOrderConnector struct {
	userID     string //the userID is a hash generate from the apikey + secretkey
	apiKey     string
	secretKey  string
	timeOffset int64
	client     *binance.Client
}

// NewOrderConnector creates a new connector that can open orders on the Binance exchange
func NewOrderConnector(apiKey, secretKey string, isTestnet bool) *BinanceOrderConnector {
	binance.UseTestnet = isTestnet
	var connector = &BinanceOrderConnector{
		apiKey:    apiKey,
		secretKey: secretKey,
		client:    binance.NewClient(apiKey, secretKey),
		userID:    hash.SHA1(apiKey + secretKey)}
	connector.timeOffset, _ = connector.client.NewSetServerTimeService().Do(context.Background())
	return connector
}

func (conn *BinanceOrderConnector) FindOrderByID(tradingPair trade.CurrencyPair, orderID string) (*trade.Order, error) {
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

func (conn *BinanceOrderConnector) NewOpenOrder(order trade.Order) (string, error) {
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

func (conn *BinanceOrderConnector) OpenOrders(tradingPair trade.CurrencyPair) ([]trade.Order, error) {
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

func (conn *BinanceOrderConnector) CancelOrder(tradingPair trade.CurrencyPair, orderID string) error {
	_, err := conn.client.NewCancelOrderService().OrderID(orderID).Symbol(tradingPair.SimpleName()).
		Do(context.Background())
	return err
}

func (conn *BinanceOrderConnector) CancelOpenOrders(tradingPair trade.CurrencyPair) error {
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
