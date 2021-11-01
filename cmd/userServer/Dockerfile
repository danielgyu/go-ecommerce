FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /dist/userServer ./cmd/userServer/main.go 

EXPOSE 8001

CMD ["/dist/userServer"]
