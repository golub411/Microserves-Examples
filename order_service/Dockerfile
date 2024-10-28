FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /order_service

EXPOSE 8002

CMD ["/order_service"]