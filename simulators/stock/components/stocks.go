package stocks

type TimeSeries struct {
	ExchangeName string
	Stocks       Stock
}

type Stock struct {
	CompanyA float64
	CompanyB float64
	CompanyC float64
	CompanyD float64
	CompanyE float64
	CompanyF float64
}
