FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /user_service

EXPOSE 8000

CMD ["/user_service"]