package trade

// InfoConnector contains information about the exchange and its products.
type InfoConnector interface {
	// GetTradingPairs returns a list of trading pairs that can be traded.
	// a not nil error indicates a connection problem with the exchange.
	TradingPairs() ([]CurrencyPair, error)

	// GetTradingRule returns the rules that must be respected in order to operate expecific trading pair
	// a not nil error indicate a connection problem with the exchange
	TradingRules(tradingPair CurrencyPair) (Rule, error)

	// CoinsBalance return the balance amount for every crypcurrency/token in the user exchange wallet
	// it is only returned assets that have a non-zero amount in either free to used or locked in some way
	// NOTE: orders placed in the orderbook are not counted in either free or locked amount
	CoinsBalance() ([]Asset, error)
}
