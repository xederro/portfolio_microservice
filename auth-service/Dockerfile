FROM golang:1.21-alpine as builder
LABEL authors="xederro"

RUN mkdir /app
COPY ./auth-service /app
COPY ./config/firebase.json /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o auth ./cmd/api
RUN chmod +x /app/auth

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/auth /app
COPY --from=builder /app/firebase.json /app

CMD [ "/app/auth" ]