FROM golang:1.23.1-alpine

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY . ./

COPY ./backend .env

RUN go build -o main ./cmd/server

EXPOSE 8080

CMD ["./main"]