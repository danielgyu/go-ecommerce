package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/danielgyu/go-ecommerce/internal/proto"
	"google.golang.org/grpc"
)

const serverAddr = ":8000"

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	ctx := context.Background()

	registerProduct(client, ctx)
	getProduct(client, ctx)
	getProducts(client, ctx)
}

func registerProduct(client pb.ProductServiceClient, ctx context.Context) {
	res, err := client.RegisterProduct(ctx, &pb.RegisterProductRequest{
		Name:  "ipad pro",
		Price: 100,
		Stock: 50,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("RegisterProduct response, id: %d\n", res.GetId())
}

func getProduct(client pb.ProductServiceClient, ctx context.Context) {
	res, err := client.GetProduct(ctx, &pb.GetProductRequest{
		Id: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("GetProduct response, id: %d, name: %s, price %d, stock: %d\n", res.Product.Id, res.Product.Name, res.Product.Price, res.Product.Stock)
}

func getProducts(client pb.ProductServiceClient, ctx context.Context) {
	res, err := client.GetProducts(ctx, &pb.GetProductsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for i, p := range res.GetProducts() {
		fmt.Printf("GetProducts response number: %d, id: %d, name: %s, price %d, stock: %d\n", i, p.Id, p.Name, p.Price, p.Stock)
	}
}
