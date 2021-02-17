go clean

set GOOS=linux

go build -a -installsuffix cgo -o service-stocktrader cmd/main.go

docker build -t vickeyshrestha/service-stock-trader:%1 .

echo "Process finished for docker build"