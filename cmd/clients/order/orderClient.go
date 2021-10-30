package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/danielgyu/go-ecommerce/internal/proto"
	"google.golang.org/grpc"
)

const serverAddr = ":8002"

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewOrderServcieClient(conn)
	ctx := context.Background()

	addToCart(client, ctx)
	removeFromCart(client, ctx)
	orderCart(client, ctx)
}

func addToCart(client pb.OrderServcieClient, ctx context.Context) {
	res, err := client.AddToCart(ctx, &pb.AddToCartRequest{
		CartId:     1,
		ProductIds: []int64{1, 2, 3},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Added %d number of items\n", res.AddedItems)
}

func removeFromCart(client pb.OrderServcieClient, ctx context.Context) {
	res, err := client.RemoveFromCart(ctx, &pb.RemoveFromCartRequest{
		CartId:    1,
		ProductId: 3,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %d number of items\n", res.Deleted)
}

func orderCart(client pb.OrderServcieClient, ctx context.Context) {
	res, err := client.OrderInCart(ctx, &pb.OrderInCartRequest{
		CartId: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ordering products:", res.ProductIds)
}
