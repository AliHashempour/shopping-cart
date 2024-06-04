FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -v -o app ./cmd/server

FROM alpine:3.12

WORKDIR /app

COPY --from=builder /app/ ./

CMD ["./app"]
