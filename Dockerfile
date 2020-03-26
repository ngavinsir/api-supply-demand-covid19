FROM golang:1.14 AS builder

COPY . /app
WORKDIR /app

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:latest AS production

WORKDIR /root/

COPY --from=builder /app .

CMD ["./app"]