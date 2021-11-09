package gateway

import (
	"context"
	"log"
	"net/http"

	pb "github.com/danielgyu/go-ecommerce/internal/proto"
	"google.golang.org/grpc"
)

type grpcClients struct {
	productClient pb.ProductServiceClient
	userClient    pb.UserServiceClient
	orderClient   pb.OrderServcieClient
}

func RunGateway() {
	c := &grpcClients{}

	registerProductClient(c)
	registerUserClient(c)
	registerOrderClient(c)

	handler := NewGatewayHandler(c)
	mux := http.NewServeMux()

	registerEndpoints(mux, handler)

	log.Println("Running gateway at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func registerEndpoints(mux *http.ServeMux, h *gatewayHandler) {
	mux.HandleFunc("/health", h.healthCheck)
	mux.HandleFunc("/initdb", h.initializedb)

	mux.HandleFunc("/userhealth/", h.userHealthCheck)
	mux.HandleFunc("/user/addcredit/", h.addCredit)
	mux.HandleFunc("/user/getcredit/", h.getCredit)
	mux.HandleFunc("/user/login/", h.logIn)
	mux.HandleFunc("/user/signup/", h.signUp)

	mux.HandleFunc("/producthealth/", h.productHealthCheck)
	mux.HandleFunc("/product/all/", h.getProductList)
	mux.HandleFunc("/product/detail/", h.getProduct)
	mux.HandleFunc("/product/new/", h.registerProduct)

	mux.HandleFunc("/orderhealth/", h.orderHealthCheck)
	mux.HandleFunc("/order/new/", h.addToCart)
	mux.HandleFunc("/order/purchase/", h.orderInCart)
	mux.HandleFunc("/order/remove", h.removeFromCart)
}

func registerProductClient(gateway *grpcClients) {
	conn, err := grpc.Dial("product-service:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewProductServiceClient(conn)

	_, err = c.HealthCheck(context.Background(), &pb.HealthCheckRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("product grpc server healthy and connected")

	gateway.productClient = c
}

func registerUserClient(gateway *grpcClients) {
	conn, err := grpc.Dial("user-service:8001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewUserServiceClient(conn)

	_, err = c.HealthCheck(context.Background(), &pb.HealthCheckRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("user grpc server healthy and connected")

	gateway.userClient = c
}

func registerOrderClient(gateway *grpcClients) {
	conn, err := grpc.Dial("order-service:8002", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewOrderServcieClient(conn)

	_, err = c.HealthCheck(context.Background(), &pb.HealthCheckRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("order grpc server healthy and connected")

	gateway.orderClient = c
}
