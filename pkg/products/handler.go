package products

import (
	"context"

	pb "github.com/danielgyu/go-ecommerce/internal/proto"
)

type productsHandler struct {
	pb.UnimplementedProductServiceServer
}

func NewProductHandler() *productsHandler {
	return &productsHandler{}
}

func (h *productsHandler) GetProduct(ctx context.Context, in pb.GetProductRequest) (out pb.GetProductResponse, err error) {
	productId := int(in.GetId())

	return
}

func (h *productsHandler) GetProducts(ctx context.Context, in pb.GetProductsRequest) (out pb.GetProductsResponse, err error) {
	return
}
