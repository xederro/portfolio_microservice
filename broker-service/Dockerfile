FROM golang:1.21-alpine as builder
LABEL authors="xederro"

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o broker ./cmd/api
RUN chmod +x /app/broker

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/broker /app

CMD [ "/app/broker" ]