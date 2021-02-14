FROM ubuntu:latest

RUN apt-get update && apt-get install -y ca-certificates && apt-get install -y apt-utils

RUN update-ca-certificates

WORKDIR /simulator/

RUN mkdir -p /simulator

COPY ./sim-stock /simulator

ENTRYPOINT ["/sim-stock"]

EXPOSE 8085