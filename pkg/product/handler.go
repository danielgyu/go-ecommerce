package product

import (
	"context"
	"database/sql"

	pb "github.com/danielgyu/go-ecommerce/internal/proto"
)

type productsHandler struct {
	pb.UnimplementedProductServiceServer
	repo *productsRepository
}

func NewProductHandler(db *sql.DB) *productsHandler {
	r := NewProductsRepository(db)
	return &productsHandler{repo: r}
}

func (h *productsHandler) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (out *pb.HealthCheckResponse, err error) {
	return &pb.HealthCheckResponse{StatusCode: 200}, nil
}

func (h *productsHandler) GetProduct(ctx context.Context, in *pb.GetProductRequest) (out *pb.GetProductResponse, err error) {
	productId := int(in.GetId())
	d, err := h.repo.GetProductDetail(productId)
	if err != nil {
		return &pb.GetProductResponse{}, err
	}

	p := &pb.Product{
		Id:    int32(d.Id),
		Name:  d.Name,
		Price: int32(d.Price),
		Stock: int32(d.Stock),
	}
	return &pb.GetProductResponse{Product: p}, nil
}

func (h *productsHandler) GetProducts(ctx context.Context, in *pb.GetProductsRequest) (out *pb.GetProductsResponse, err error) {
	ps, err := h.repo.GetProductList()
	if err != nil {
		return &pb.GetProductsResponse{}, err
	}

	products := make([]*pb.Product, len(ps))
	for i, p := range ps {
		products[i] = &pb.Product{
			Id:    int32(p.Id),
			Name:  p.Name,
			Price: int32(p.Price),
			Stock: int32(p.Stock),
		}
	}
	return &pb.GetProductsResponse{Products: products}, nil
}

func (h *productsHandler) RegisterProduct(ctx context.Context, in *pb.RegisterProductRequest) (out *pb.RegisterProductResponse, err error) {
	id, err := h.repo.RegisterProduct(in.GetName(), int(in.GetPrice()), int(in.GetStock()))
	if err != nil {
		return &pb.RegisterProductResponse{}, err
	}

	return &pb.RegisterProductResponse{Id: int32(id)}, nil
}
