FROM golang:1.21.6-alpine as builder

WORKDIR /app

COPY ./ ./

RUN mkdir bin

RUN go mod tidy

RUN go build -o ./bin/cockroachApp ./main.go

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/config.json ./
COPY --from=builder /app/bin/cockroachApp ./

EXPOSE 5000
CMD ./cockroachApp

