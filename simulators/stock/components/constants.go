package stocks

// error constants
const (
	ErrFailedToInitializeNatsClient = "error while initializing nats client %s"
	ErrNatsHostPortRequired         = "error - nats host and port number are required as env variables"
)

// info constants
const (
	InfoNatsConnection = "nats successfully connected to %s"
)

// nats
const (
	SimNatsTopic = "SIM.TS.Stocks"
)
