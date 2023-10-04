FROM golang:1.21-alpine as builder
LABEL authors="xederro"

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o weather ./cmd/api
RUN chmod +x /app/weather

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/weather /app

CMD [ "/app/weather" ]