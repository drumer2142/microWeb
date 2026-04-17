FROM golang:1.21-alpine

RUN mkdir -p /app

WORKDIR /app

COPY ./src /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

CMD ["./app"]
