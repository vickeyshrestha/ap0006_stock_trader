package stocktrader

//var ApplicationConfiguration configuration

type Configuration struct {
	GrpcPort     string
	Httpport     string
	ServiceName  string
	DatabaseHost string
	DatabasePort int
	Environment  string
}
