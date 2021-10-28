package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	pb "github.com/danielgyu/go-ecommerce/internal/proto"
	"google.golang.org/grpc"
)

const userAddr = ":8001"

func main() {
	conn, err := grpc.Dial(userAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	ctx := context.Background()

	u := signUpUser(client, ctx)
	t := logInUser(client, ctx, u)
	addCredit(client, ctx, t)
}

func signUpUser(client pb.UserServiceClient, ctx context.Context) string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10000)
	u := fmt.Sprintf("user%d", r)
	res, err := client.SignUp(ctx, &pb.SignUpRequest{
		Username: u,
		Password: "password",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("SignUpUser response status: %t\n", res.Success)
	return u
}

func logInUser(client pb.UserServiceClient, ctx context.Context, username string) string {
	res, err := client.LogIn(ctx, &pb.LogInRequest{
		Username: username,
		Password: "password",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("LogInUser reponse, token: %s\n", res.Token)
	return res.Token
}

func addCredit(client pb.UserServiceClient, ctx context.Context, token string) {
	res, err := client.AddCredit(ctx, &pb.AddCreditRequest{
		Token:  token,
		Credit: 99,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("AddCredit response, new credit is %d\n", res.Credit)
}
