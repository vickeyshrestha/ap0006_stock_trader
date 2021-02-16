module github/godzilla/services/stock-trader

go 1.15

replace github.com/vickeyshrestha/sharing-services/protobuf/stock_trader => ../../../github.com/vickeyshrestha/sharing-services/protobuf/stock_trader

replace github.com/vickeyshrestha/sharing-services/drivers/postgres => ../../../github.com/vickeyshrestha/sharing-services/drivers/postgres

require (
	github.com/labstack/echo/v4 v4.2.0
	//github.com/vickeyshrestha/sharing-services v0.0.6 // indirect
	github.com/vickeyshrestha/sharing-services/drivers/postgres v0.0.5
	github.com/vickeyshrestha/sharing-services/protobuf/stock_trader v0.0.5
	google.golang.org/protobuf v1.25.0
)
