FROM ubuntu:latest

RUN apt-get update && apt-get install -y ca-certificates && apt-get install -y apt-utils

RUN update-ca-certificates

WORKDIR /service/

COPY ./stock-trader /service

ENTRYPOINT ["./stock-trader"]

EXPOSE 8085