package order

import (
	"context"
	"database/sql"
	"errors"
)

type orderRepository struct {
	database *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{database: db}
}

func (r *orderRepository) PutIntoCart(ctx context.Context, cartId int64, productIds []int64) (affected int64, err error) {
	var InsertToCart string = "INSERT INTO cart_products (cart_id, product_id) VALUES (?, ?)"

	tx, err := r.database.BeginTx(ctx, nil)
	if err != nil {
		return 0, errors.New("db failure")
	}

	defer tx.Rollback()

	for _, p := range productIds {
		res, err := r.database.Exec(InsertToCart, cartId, p)
		if err != nil {
			return 0, err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			return 0, err
		}

		affected += rows
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return affected, nil
}

func (r *orderRepository) DeleteInCart(ctx context.Context, cartId int64, productId int64) (deleted int64, err error) {
	var DeleteProduct = "DELETE FROm cart_products WHERE cart_id = ? and product_id = ?"

	res, err := r.database.Exec(DeleteProduct, cartId, productId)
	if err != nil {
		return 0, err
	}

	deleted, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}

func (r *orderRepository) GetAllCartProducts(ctx context.Context, cartId int64) (productList []int64, err error) {
	var GetProductsList = "SELECT product_id FROM cart_products WHERE cart_id = ?"

	rows, err := r.database.Query(GetProductsList, cartId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product_id int64
		if err = rows.Scan(&product_id); err != nil {
			return nil, err
		}
		productList = append(productList, product_id)
	}

	return productList, nil
}
