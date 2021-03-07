package stocks

type TimeSeries struct {
	ExchangeName string
	Stocks       Stock
}

type Stock struct {
	Companies map[string]float64
}
