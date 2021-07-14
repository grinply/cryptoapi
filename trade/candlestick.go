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

type CandleInterval struct {
	description string //description of the interval = [1m, 3m, 5m, 15m, 30m, 1h, 4h, 1D]
	minutes     int64
}

// Minutes returns the amount of minutes in this interval
func (interval CandleInterval) Minutes() int64 {
	return interval.minutes
}

// Description returns the string containing a tag description for the candle Interval
// examples: 1m, 3m, 5m, 15m, 30m, 1h, 4h, 12h, 1D
func (interval CandleInterval) Description() string {
	return interval.description
}

var (
	// OneMinute represents the 1 minute interval between between candlestick with price information
	// candlesticks with interval one minute will agregate all prices in this timeframe
	OneMinute CandleInterval = CandleInterval{description: "1m", minutes: 1}

	// OneMinute represents the 3 minutes interval between between candlestick with price information
	// candlesticks with interval one minute will agregate all prices in this timeframe
	ThreeMinutes CandleInterval = CandleInterval{description: "3m", minutes: 3}

	// OneMinute represents the 5 minutes interval between between candlestick with price information
	// candlesticks with interval one minute will agregate all prices in this timeframe
	FiveMinutes CandleInterval = CandleInterval{description: "5m", minutes: 5}

	// OneMinute represents the 15 minutes interval between between candlestick with price information
	// candlesticks with interval one minute will agregate all prices in this timeframe
	FifteenMinutes CandleInterval = CandleInterval{description: "15m", minutes: 15}

	// OneMinute represents the 30 minutes interval between between candlestick with price information
	// candlesticks with interval one minute will agregate all prices in this timeframe
	ThirtyMinutes CandleInterval = CandleInterval{description: "30m", minutes: 30}

	// OneMinute represents the 1 hour interval between between candlestick with price information
	// candlesticks with interval one minute will agregate all prices in this timeframe
	OneHour CandleInterval = CandleInterval{description: "1h", minutes: 60}

	// OneMinute represents the 4 hour interval between between candlestick with price information
	// candlesticks with interval one minute will agregate all prices in this timeframe
	FourHours CandleInterval = CandleInterval{description: "4h", minutes: 240}

	// OneMinute represents the 12 hour interval between between candlestick with price information
	// candlesticks with interval one minute will agregate all prices in this timeframe
	TwelveHours CandleInterval = CandleInterval{description: "12h", minutes: 720}

	// OneDay represents the 1 day interval between between candlestick with price information
	// candlesticks with interval one minute will agregate all prices in this timeframe
	OneDay CandleInterval = CandleInterval{description: "1D", minutes: 1440}
)

type Rule struct {
	TradingPair         CurrencyPair
	BaseAssetPrecision  int32
	QuoteAssetPrecision int32
	MaxOrderQty         int    // maximum amount of orders that can be present in the exchange orderbook simultaneously
	MinPriceMovement    string // the minimum price variation that can occur in the trading pair
	BaseStepSize        string // defines the step multiplier that can be applied for the trading pair, any amount of base currency MUST be a multiple of this value
	MinNotionalValue    string // is the minimum required notional value to open a order in the exchange, calculate by: price * amount
	MinLotSize          string // is the minimum value in quote currency that a order must have in order to be accepted by the exchange
}
