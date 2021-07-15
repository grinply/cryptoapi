package trade

// CandleStick represents the pricing information for a given TradingPair, its the format commonly used in crypto trading
// for represent price data.
//
// Open is the starting price of the candlestick
// High is the maximum price reached in the timeframe of the candlestick
// Low is the lowest price reached in the timeframe of the candlestick
// Close is the last price before the candlestick has finished
// Volume is the total amount that has been traded in the interval while the candlestick is active
type CandleStick struct {
	TradingPair CurrencyPair
	Open        string `json:"open_price"`
	High        string `json:"high_price"`
	Low         string `json:"low_price"`
	Close       string `json:"close_price"`
	Volume      string `json:"volume"`
	StartTime   string `json:"start_time"`
	IsClosed    bool   `json:"is_closed"`
}

// Rule is a description of requirements that exchanges impose in order to trade a given currency pair
type Rule struct {
	TradingPair         CurrencyPair
	BaseAssetPrecision  int
	QuoteAssetPrecision int
	MaxOrderQty         int    // maximum amount of orders that can be present in the exchange orderbook simultaneously
	MinPriceMovement    string // the minimum price variation that can occur in the trading pair
	BaseStepSize        string // defines the step multiplier that can be applied for the trading pair, any amount of base currency MUST be a multiple of this value
	MinNotionalValue    string // is the minimum required notional value to open a order in the exchange, calculate by: price * amount
	MinLotSize          string // is the minimum value in quote currency that a order must have in order to be accepted by the exchange
}
