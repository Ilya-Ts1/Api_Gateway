FROM golang:1.23.0-latest AS builder

WORKDIR /api_gateway

COPY . .

RUN go mod download

COPY . .

RUN go build -o api_gateway ./cmd/app

FROM alpine:latest
WORKDIR /api_gateway

COPY --from=builder /api_gateway/api_gateway .

EXPOSE 8080
CMD ["./api_gateway"]