:: This batch file downloads and installs the NATS server and then starts the NATS server by running the executable gnatsd
ECHO starting NATS server
go get github.com/nats-io/gnatsd
gnatsd
:: docker run -d --name nats-main -p 4222:4222 -p 6222:6222 -p 8222:8222 nats