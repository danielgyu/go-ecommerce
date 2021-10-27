package product

import (
	"database/sql"
)

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

func (r *productsRepository) GetProductDetail(productId int) (p ProductDetail, err error) {
	var QueryOneProduct string = "SELECT * FROM products WHERE id = ?"

	if err = r.database.QueryRow(QueryOneProduct, productId).Scan(&p.Id, &p.Name, &p.Price, &p.Stock); err != nil {
		return p, err
	}

	return
}

func (r *productsRepository) GetProductList() (ps []ProductDetail, err error) {
	var QueryAllProducts string = "SELECT * FROM products"

	rows, err := r.database.Query(QueryAllProducts)
	if err != nil {
		return ps, err
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return ps, err
	}

	for rows.Next() {
		p := new(ProductDetail)
		if err := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Stock); err != nil {
			return ps, err
		}

		ps = append(ps, *p)
	}

	return
}

func (r *productsRepository) RegisterProduct(name string, price int, stock int) (id int64, err error) {
	var InsertProduct string = "INSERT INTO products (name, price, stock) VALUES (?, ?, ?)"

	result, err := r.database.Exec(InsertProduct, name, price, stock)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return
}
