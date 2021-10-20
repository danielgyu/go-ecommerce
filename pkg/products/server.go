package products

import (
	"fmt"
	"log"
	"net"

	pb "github.com/danielgyu/go-ecommerce/internal/proto"

	"google.golang.org/grpc"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	address := fmt.Sprintf("localhost:%d", 9001)
	lis, err := net.Listen("tcp", address)
	check(err)

	grpcProducts := grpc.NewServer()
	ei := NewProductHandler()
	pb.RegisterProductServiceServer(grpcProducts, ei)

	grpcProducts.Serve(lis)
}
