package trade

// PriceConnector enables communication with a cryptocurrency exchange in order to retrieve pricing and ot information and
type PriceConnector interface {
	// GetLatestPrice returns the lastest and most updated price for the provided trading pair in string format
	// a not nil error indicate a connection problem with the exchange
	GetLatestPrice(tradingPair CurrencyPair) (string, error)

	// GetTradingRule returns the rules that must be respected in order to operate expecific trading pair
	// a not nil error indicate a connection problem with the exchange
	GetTradingRule(tradingPair CurrencyPair) (*Rule, error)

	//GetCandles returns the most recent candles available for the provided tradingPair in the desired interval
	GetCandles(tradingPair CurrencyPair, qty int, interval CandleTime) ([]CandleStick, error)

	// GetPriceFeed returns a channel that can be used to consume CandleStick pricing data in realtime
	// the IsClosed parameter of each CandleStick will be true when a timeinterval of 1 minute has passed for the candlestick
	GetPriceFeed(tradingPair CurrencyPair) <-chan CandleStick
}
