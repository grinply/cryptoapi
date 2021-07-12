package trade

// OrderConnector enables communication with a cryptocurrency exchange to execute trading order.
// Trading orders are actions that allows a user to buy or sell amounts of a given cryptocurrency.
// For example: buy 12 ADA in the ADA/ETH pair.
//
// Note: all the communication using the OrderConnector is for single user only, it is required to authenticated providing api keys.
type OrderConnector interface {

	// GetWalletBalances checks the balance amount for every crypcurrency/token in the user exchange wallet
	GetWalletBalances() ([]Asset, error)

	// GetOrderByID returns all the information about a order that contains the provided ID
	// a not nill error indicate a connection problem with the exchange
	// when the error is nil and the *Order is nil it means that no order could be found on exchange for the provided ID
	GetOrderByID(tradingPair CurrencyPair, orderID string) (*Order, error)

	// GetAllOpenOrders returns a slice containing all open Orders (present in the orderbook) for the provided trading pair
	// a not nill error indicate a connection problem with the exchange
	GetAllOpenOrders(tradingPair CurrencyPair) ([]Order, error)

	// OpenOrder executes buys or sells a amount of base currency in some price on the exchange.
	// If the order type is 'market' the order will be executed in the best avaliable price in the orderbook of the exchange.
	// If the order type is 'limit' the order will be placed in the exchange orderbook and only executed when the desired price is reached.
	// a not nil error indicate a connection problem with the exchange
	OpenOrder(orderToOpen Order) (string, error)

	// CancelOrder removes a order from the exchange order book if the order wasn't already executed
	// a nil error return indicates the orders was sucessfuly closed or it had been previosly executed
	// a not nil error indicate the order was not found in the for the provided id OR a connection problem with the exchange
	CancelOrder(tradingPair CurrencyPair, orderID string) error

	// CancelAllOpenOrders cancels all open orders that were previously placed in the exchange orderbook for the provided trading pair.
	// a not nil error indicate a connection problem with the exchange, a nil value denotes that all orders in the orderbook were canceled sucessfuly.
	CancelAllOpenOrders(tradingPair CurrencyPair) error
}
