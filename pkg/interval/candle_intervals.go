package interval

import "github.com/grinply/cryptoapi/pkg/trade"

var (
	// OneMinute represents the 1 minute interval between between trade.CandleTimestick with price information
	// trade.CandleTimesticks with interval one minute will agregate all prices in this timeframe
	OneMinute trade.CandleTime = trade.NewCandleTime("1m", 1)

	// OneMinute represents the 3 minutes interval between between trade.CandleTimestick with price information
	// trade.CandleTimesticks with interval one minute will agregate all prices in this timeframe
	ThreeMinutes trade.CandleTime = trade.NewCandleTime("3m", 3)

	// OneMinute represents the 5 minutes interval between between trade.CandleTimestick with price information
	// trade.CandleTimesticks with interval one minute will agregate all prices in this timeframe
	FiveMinutes trade.CandleTime = trade.NewCandleTime("5m", 5)

	// OneMinute represents the 15 minutes interval between between trade.CandleTimestick with price information
	// trade.CandleTimesticks with interval one minute will agregate all prices in this timeframe
	FifteenMinutes trade.CandleTime = trade.NewCandleTime("15m", 15)

	// OneMinute represents the 30 minutes interval between between trade.CandleTimestick with price information
	// trade.CandleTimesticks with interval one minute will agregate all prices in this timeframe
	ThirtyMinutes trade.CandleTime = trade.NewCandleTime("30m", 30)

	// OneMinute represents the 1 hour interval between between trade.CandleTimestick with price information
	// trade.CandleTimesticks with interval one minute will agregate all prices in this timeframe
	OneHour trade.CandleTime = trade.NewCandleTime("1h", 60)

	// OneMinute represents the 4 hour interval between between trade.CandleTimestick with price information
	// trade.CandleTimesticks with interval one minute will agregate all prices in this timeframe
	FourHours trade.CandleTime = trade.NewCandleTime("4h", 240)

	// OneMinute represents the 12 hour interval between between trade.CandleTimestick with price information
	// trade.CandleTimesticks with interval one minute will agregate all prices in this timeframe
	TwelveHours trade.CandleTime = trade.NewCandleTime("12h", 720)

	// OneDay represents the 1 day interval between between trade.CandleTimestick with price information
	// trade.CandleTimesticks with interval one minute will agregate all prices in this timeframe
	OneDay trade.CandleTime = trade.NewCandleTime("1D", 1440)
)
