FROM golang:1.14.2-alpine3.11 as compile

WORKDIR /app

FROM alpine:3.11
COPY --from=compile /app /app
