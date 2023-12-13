package stocktrader

// error constants
const (
	ErrConnectingToDb = "error while connecting to database %s. Retrying connection attempt %d"
)

// info const
const (
	InfoSuccessfulConnectionDb = "successfully connected to database"
)

// general constants
const (
	serviceName                     = "stock-traders"
	statusAlive                     = "Alive"
	SimNatsTopic                    = "SIM.TS.Stocks"
	DequeueProcessIntervalInSeconds = 8
	TimeToCalculateValuesAndDequeue = 8
	QueueSize                       = 15
)
