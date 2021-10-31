package gateway

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	pb "github.com/danielgyu/go-ecommerce/internal/proto"
	"github.com/golang/protobuf/jsonpb"
)

func (h *gatewayHandler) productHealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	_, err := h.clients.productClient.HealthCheck(ctx, &pb.HealthCheckRequest{})
	if err != nil {
		log.Println("error healthchecking user grpc")
		w.Write([]byte("error"))
		return
	}

	w.Write([]byte("user grpc healthy"))
}

func (h *gatewayHandler) getProductList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := pb.GetProductsRequest{}

	ctx := context.Background()
	res, err := h.clients.productClient.GetProducts(ctx, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *gatewayHandler) getProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/product/detail/"))
	if err != nil {
		errorResponse(err, w)
		return
	}

	req := pb.GetProductRequest{Id: int32(id)}
	ctx := context.Background()
	res, err := h.clients.productClient.GetProduct(ctx, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *gatewayHandler) registerProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := pb.RegisterProductRequest{}
	err := jsonpb.Unmarshal(r.Body, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	ctx := context.Background()
	res, err := h.clients.productClient.RegisterProduct(ctx, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	json.NewEncoder(w).Encode(res)
}
