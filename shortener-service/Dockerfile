FROM golang:1.21-alpine as builder
LABEL authors="xederro"

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o shortener ./cmd/api
RUN chmod +x /app/shortener

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/shortener /app

CMD [ "/app/shortener" ]