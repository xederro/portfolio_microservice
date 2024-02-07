FROM golang:1.21-alpine as builder
LABEL authors="xederro"

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o frontend ./cmd/web
RUN chmod +x /app/frontend

FROM alpine:latest
RUN mkdir /app
RUN mkdir /app/cmd
RUN mkdir /app/cmd/web
RUN mkdir /app/cmd/web/data
RUN mkdir /app/cmd/web/public
RUN mkdir /app/cmd/web/templates
COPY --from=builder /app/frontend /app
COPY --from=builder /app/cmd/web/templates /cmd/web/templates
COPY --from=builder /app/cmd/web/public /cmd/web/public
COPY --from=builder /app/cmd/web/data /cmd/web/data

CMD [ "/app/frontend" ]