package stocktrader

var ApplicationConfiguration configuration

type configuration struct {
	GrpcPort     string
	Httpport     string
	ServiceName  string
	DatabaseHost string
	databasePort int
	Environment  string
}
