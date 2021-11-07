FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /dist/gateway ./cmd/gateway/main.go 
RUN apk --no-cache add curl

EXPOSE 8080

CMD ["/app/wait_for_servers.sh", "/dist/gateway"]
