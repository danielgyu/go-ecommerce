package gateway

import (
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

	mux.HandleFunc("/userhealth", h.userHealthCheck)
	mux.HandleFunc("/user/addcredit", h.addCredit)
	mux.HandleFunc("/user/login", h.logIn)
	mux.HandleFunc("/user/signup", h.signUp)
}

func registerProductClient(gateway *grpcClients) {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewProductServiceClient(conn)
	gateway.productClient = c
}

func registerUserClient(gateway *grpcClients) {
	conn, err := grpc.Dial(":8001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewUserServiceClient(conn)
	gateway.userClient = c
}

func registerOrderClient(gateway *grpcClients) {
	conn, err := grpc.Dial(":8002", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewOrderServcieClient(conn)
	gateway.orderClient = c
}
