package order

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/danielgyu/go-ecommerce/internal/proto"
)

type orderHandler struct {
	pb.UnimplementedOrderServcieServer
	repo *orderRepository
}

func NewOrderHandler(db *sql.DB) *orderHandler {
	r := NewOrderRepository(db)
	return &orderHandler{repo: r}
}

func (h *orderHandler) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (out *pb.HealthCheckResponse, err error) {
	return &pb.HealthCheckResponse{StatusCode: 200}, nil
}

func (h *orderHandler) AddToCart(ctx context.Context, in *pb.AddToCartRequest) (out *pb.AddToCartResponse, err error) {
	added, err := h.repo.PutIntoCart(ctx, in.CartId, in.ProductIds)
	if err != nil {
		return &pb.AddToCartResponse{}, err
	}

	return &pb.AddToCartResponse{AddedItems: int64(added)}, nil
}

func (h *orderHandler) RemoveFromCart(ctx context.Context, in *pb.RemoveFromCartRequest) (out *pb.RemoveFromCartResponse, err error) {
	deleted, err := h.repo.DeleteInCart(ctx, in.CartId, in.ProductId)
	if err != nil {
		log.Println("error while removing from cart")
		return &pb.RemoveFromCartResponse{}, err
	}

	return &pb.RemoveFromCartResponse{Deleted: deleted}, nil
}

func (h *orderHandler) OrderInCart(ctx context.Context, in *pb.OrderInCartRequest) (out *pb.OrderInCartResponse, err error) {
	productList, err := h.repo.GetAllCartProducts(ctx, in.UserId)
	if err != nil {
		log.Println("error retrieving product list")
		return &pb.OrderInCartResponse{}, err
	}

	return &pb.OrderInCartResponse{ProductIds: productList}, nil
}
