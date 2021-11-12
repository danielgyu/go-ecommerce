package gateway

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	pb "github.com/danielgyu/go-ecommerce/internal/proto"
	"github.com/golang/protobuf/jsonpb"
)

func (h *gatewayHandler) orderHealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	_, err := h.clients.orderClient.HealthCheck(ctx, &pb.HealthCheckRequest{})
	if err != nil {
		log.Println("error healthchecking user grpc")
		w.Write([]byte("error"))
		return
	}

	w.Write([]byte("user grpc healthy"))
}

func (h *gatewayHandler) addToCart(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.Background()

	type AddToCart struct {
		Token      string  `json:"token"`
		ProductIds []int64 `json:"productIds"`
	}

	req := AddToCart{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorResponse(err, w)
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	uReq := pb.GetUserIdRequest{Token: req.Token}
	uRes, err := h.clients.userClient.GetUserId(ctx, &uReq)
	log.Println("Get userid response:")
	log.Println(uRes)

	cReq := pb.AddToCartRequest{UserId: uRes.UserId, ProductIds: req.ProductIds}
	res, err := h.clients.orderClient.AddToCart(ctx, &cReq)
	if err != nil {
		errorResponse(err, w)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *gatewayHandler) removeFromCart(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := pb.RemoveFromCartRequest{}
	err := jsonpb.Unmarshal(r.Body, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	ctx := context.Background()
	res, err := h.clients.orderClient.RemoveFromCart(ctx, &req)
	if err != nil {
		errorResponse(err, w)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *gatewayHandler) orderInCart(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.Background()

	t := pb.GetCreditRequest{}
	err := jsonpb.Unmarshal(r.Body, &t)
	if err != nil {
		errorResponse(err, w)
	}
	creditRes, err := h.clients.userClient.GetCredit(ctx, &t)
	if err != nil {
		errorResponse(err, w)
	}

	orderReq := pb.OrderInCartRequest{UserId: creditRes.UserId}
	orderRes, err := h.clients.orderClient.OrderInCart(ctx, &orderReq)
	if err != nil {
		errorResponse(err, w)
		return
	}

	var priceAggregate int64 = 0
	productIds := orderRes.GetProductIds()
	for _, p := range productIds {
		r := pb.GetProductRequest{Id: int32(p)}
		res, err := h.clients.productClient.GetProduct(ctx, &r)
		if err != nil {
			errorResponse(err, w)
			priceAggregate += int64(res.Product.Price)
		}
	}

	if priceAggregate > creditRes.Credit {
		w.Write([]byte("not enough credit"))
	}
	w.Write([]byte("enough credit, proceed to payment"))
}
