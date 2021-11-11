package gateway

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pb "github.com/danielgyu/go-ecommerce/internal/proto"
	"github.com/golang/protobuf/jsonpb"
)

func (h *gatewayHandler) userHealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	_, err := h.clients.userClient.HealthCheck(ctx, &pb.HealthCheckRequest{})
	if err != nil {
		log.Println("error healthchecking user grpc")
		w.Write([]byte("error"))
		return
	}

	w.Write([]byte("user grpc healthy"))
}

func (h *gatewayHandler) signUp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := pb.SignUpRequest{}
	err := jsonpb.Unmarshal(r.Body, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	ctx := context.Background()
	res, err := h.clients.userClient.SignUp(ctx, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	oReq := pb.RegisterUserCartRequest{UserId: res.UserId}
	oRes, err := h.clients.orderClient.RegisterUserCart(ctx, &oReq)
	if err != nil {
		errorResponse(err, w)
		return
	}

	json.NewEncoder(w).Encode(oRes)
}

func (h *gatewayHandler) logIn(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := pb.LogInRequest{}
	err := jsonpb.Unmarshal(r.Body, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	ctx := context.Background()
	res, err := h.clients.userClient.LogIn(ctx, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *gatewayHandler) getCredit(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := pb.GetCreditRequest{}
	err := jsonpb.Unmarshal(r.Body, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	ctx := context.Background()
	res, err := h.clients.userClient.GetCredit(ctx, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *gatewayHandler) addCredit(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := pb.AddCreditRequest{}
	err := jsonpb.Unmarshal(r.Body, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	ctx := context.Background()
	res, err := h.clients.userClient.AddCredit(ctx, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	json.NewEncoder(w).Encode(res)
}
