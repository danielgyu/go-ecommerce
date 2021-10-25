package products

import "database/sql"

type productsRepository struct {
	database *sql.DB
}

func NewProductsRepository(db *sql.DB) *productsRepository {
	return &productsRepository{database: db}
}

type ProductDetail struct {
	Id    int
	Name  string
	Price int
	Stock int
}

func (r *productsRepository) GetProductDetail(productId int) (product ProductDetail, err error) {
	return
}

func (r *productsRepository) GetProductList() (products []ProductDetail, err error) {
	return
}
