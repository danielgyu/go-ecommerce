FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /dist/orderServer ./cmd/orderServer/main.go 

EXPOSE 8002

CMD ["/dist/orderServer"]
