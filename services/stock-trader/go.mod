module github/maxzilla/services/stock-trader

go 1.15

replace github.com/vickeyshrestha/sharing-services/protobuf/stock_trader => ../../../github.com/vickeyshrestha/sharing-services/protobuf/stock_trader

replace github.com/vickeyshrestha/sharing-services/drivers/sql => ../../../github.com/vickeyshrestha/sharing-services/drivers/sql

require (
	//github.com/vickeyshrestha/sharing-services v0.0.6 // indirect
	github.com/vickeyshrestha/sharing-services/drivers/sql v0.0.4
	github.com/vickeyshrestha/sharing-services/protobuf/stock_trader v0.0.4
	google.golang.org/protobuf v1.25.0
)
