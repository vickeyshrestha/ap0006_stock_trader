module github/godzilla/services/stock-trader

go 1.15

replace github.com/vickeyshrestha/sharing-services/protobuf/stock_trader => ../../../github.com/vickeyshrestha/sharing-services/protobuf/stock_trader
replace github.com/vickeyshrestha/sharing-services/drivers/postgres => ../../../github.com/vickeyshrestha/sharing-services/drivers/postgres
replace github.com/vickeyshrestha/sharing-services/drivers/nats => ../../../github.com/vickeyshrestha/sharing-services/drivers/nats

require (
	github.com/gorilla/handlers v1.5.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.1.0
	github.com/jnewmano/grpc-json-proxy v0.0.2
	github.com/vickeyshrestha/sharing-services/drivers/postgres v0.0.7
	github.com/vickeyshrestha/sharing-services/drivers/nats v0.0.7
	github.com/vickeyshrestha/sharing-services/protobuf/stock_trader v0.0.7
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
