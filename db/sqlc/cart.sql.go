// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: cart.sql

package db

import (
	"context"
)

const createCart = `-- name: CreateCart :one
INSERT into shopping_carts (
    cart_id,
    customer_id, 
    product_id
) VALUES (
    $1, $2, $3
) RETURNING cart_id, customer_id, product_id, created_at
`

type CreateCartParams struct {
	CartID     string `json:"cart_id"`
	CustomerID string `json:"customer_id"`
	ProductID  string `json:"product_id"`
}

func (q *Queries) CreateCart(ctx context.Context, arg CreateCartParams) (ShoppingCart, error) {
	row := q.db.QueryRowContext(ctx, createCart, arg.CartID, arg.CustomerID, arg.ProductID)
	var i ShoppingCart
	err := row.Scan(
		&i.CartID,
		&i.CustomerID,
		&i.ProductID,
		&i.CreatedAt,
	)
	return i, err
}

const findCart = `-- name: FindCart :many
SELECT shopping_carts.cart_id, customers.customer_name, products.product_name FROM shopping_carts
    INNER JOIN customers ON shopping_carts.customer_id=customers.customer_id
    INNER JOIN products on shopping_carts.product_id=products.product_id
    WHERE customers.customer_name=$1
`

type FindCartRow struct {
	CartID       string `json:"cart_id"`
	CustomerName string `json:"customer_name"`
	ProductName  string `json:"product_name"`
}

func (q *Queries) FindCart(ctx context.Context, customerName string) ([]FindCartRow, error) {
	rows, err := q.db.QueryContext(ctx, findCart, customerName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindCartRow
	for rows.Next() {
		var i FindCartRow
		if err := rows.Scan(&i.CartID, &i.CustomerName, &i.ProductName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}