go clean

set GOOS=linux

go build -a -installsuffix cgo -o simulator-stock cmd/main.go

docker build -t vickeyshrestha/simulator-stock:%1 .

echo "Process finished for docker build"