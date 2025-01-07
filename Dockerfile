FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main ./cmd/ordersystem

RUN cp cmd/ordersystem/.env .env

CMD [ "./main" ]
