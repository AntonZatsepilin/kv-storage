FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./

RUN apk add --no-cache openssl-dev gcc musl-dev
RUN go build -tags openssl -o main ./cmd/server

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/main .
COPY backend/configs/config.yaml ./configs/

EXPOSE 8080

CMD ["./main"]