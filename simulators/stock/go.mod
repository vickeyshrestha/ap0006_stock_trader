module godzilla/simulators/stock

go 1.15

replace github.com/vickeyshrestha/sharing-services/drivers/nats => ../../../github.com/vickeyshrestha/sharing-services/drivers/nats

require (
	//github.com/vickeyshrestha/sharing-services v0.0.6 // indirect
	github.com/vickeyshrestha/sharing-services/drivers/nats v0.0.4
	google.golang.org/protobuf v1.25.0
)
