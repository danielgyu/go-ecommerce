FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /dist/productServer ./cmd/productServer/main.go 

EXPOSE 8000

CMD ["/dist/productServer"]
