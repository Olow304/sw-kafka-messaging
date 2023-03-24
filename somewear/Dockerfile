FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/internal
RUN go build -o main

EXPOSE 8080

CMD ["/app/internal/main"]
