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

func (r *orderRepository) AddNewCart(ctx context.Context, userId int64) (added bool, err error) {
	var InsertNewCart string = "INSERT INTO carts (user_id) VALUES (?)"

	result, err := r.database.Exec(InsertNewCart, userId)
	if err != nil {
		return false, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	} else if affected == 0 {
		return false, err
	}

	return true, nil
}

func (r *orderRepository) PutIntoCart(ctx context.Context, userId int64, productIds []int64) (affected int64, err error) {
	var InsertToCart string = "INSERT INTO cart_products (user_id, product_id) VALUES (?, ?)"

	tx, err := r.database.BeginTx(ctx, nil)
	if err != nil {
		return 0, errors.New("db failure")
	}

	defer tx.Rollback()

	for _, p := range productIds {
		res, err := r.database.Exec(InsertToCart, userId, p)
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

func (r *orderRepository) DeleteInCart(ctx context.Context, userId int64, productId int64) (deleted int64, err error) {
	var DeleteProduct = "DELETE FROM cart_products WHERE user_id = ? and product_id = ?"

	res, err := r.database.Exec(DeleteProduct, userId, productId)
	if err != nil {
		return 0, err
	}

	deleted, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}

func (r *orderRepository) GetAllCartProducts(ctx context.Context, userId int64) (productList []int64, err error) {
	var GetProductList = "SELECT product_id FROM cart_products WHERE user_id = ?"

	rows, err := r.database.Query(GetProductList, userId)
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
