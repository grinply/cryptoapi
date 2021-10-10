package trade

type CandleTime struct {
	description string
	minutes     int64
}

// Minutes returns the amount of minutes in this interval
func (interval CandleTime) Minutes() int64 {
	return interval.minutes
}

// Description returns the string containing a tag description for the CandleTime Interval
// examples: 1m, 3m, 5m, 15m, 30m, 1h, 4h, 12h, 1D
func (interval CandleTime) Description() string {
	return interval.description
}

func NewCandleTime(description string, minutes int64) CandleTime {
	return CandleTime{
		description: description,
		minutes:     minutes,
	}
}
