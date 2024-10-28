FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /product_service

EXPOSE 8001

CMD ["/product_service"]