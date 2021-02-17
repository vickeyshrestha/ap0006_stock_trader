FROM alpine:latest

WORKDIR /service/
RUN mkdir /service/resources
COPY ./resources/config.json /service
COPY ./service-stocktrader /service

ENTRYPOINT ["./service-stocktrader"]

EXPOSE 8085