:: This batch file downloads and installs the NATS server and then starts the NATS server by running the executable gnatsd
ECHO starting NATS server
go get github.com/nats-io/gnatsd
gnatsd