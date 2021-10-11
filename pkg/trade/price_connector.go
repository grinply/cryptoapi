package trade

// PriceConnector enables communication with a cryptocurrency exchange in order to retrieve pricing and ot information and
type PriceConnector interface {
	// LatestPrice returns the lastest and most updated price for the provided trading pair in string format.
	// a not nil error indicate a connection problem with the exchange.
	LatestPrice(tradingPair CurrencyPair) (string, error)

	// Candles returns the most recent candles available for the provided tradingPair in the desired interval.
	// qty is the number of candles to return and interval is the time between each candle.
	// a not nil error indicate a connection problem with the exchange
	Candles(tradingPair CurrencyPair, qty int, interval CandleTime) ([]CandleStick, error)

	// PriceFeed returns a channel that can be used to consume CandleStick pricing data in realtime.
	// the IsClosed parameter of each CandleStick will be true when a time interval of 1 minute has passe.
	// each candlestick denotes up to 1 minute of price information.
	PriceFeed(tradingPair CurrencyPair) <-chan CandleStick
}
