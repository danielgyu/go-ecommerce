package products

import pb "github.com/danielgyu/go-ecommerce/internal/proto"

type productsHandler struct {
	pb.UnimplementedProductServiceServer
}

func NewProductHandler() *productsHandler {
	return &productsHandler{}
}
