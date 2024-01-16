// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: product.sql

package db

import (
	"context"
	"time"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
    product_id,
    category_id,
    product_name,
    price
) VALUES (
    $1, $2, $3, $4
) RETURNING product_id, category_id, product_name, price, created_at
`

type CreateProductParams struct {
	ProductID   string `json:"product_id"`
	CategoryID  int64  `json:"category_id"`
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.ProductID,
		arg.CategoryID,
		arg.ProductName,
		arg.Price,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.CategoryID,
		&i.ProductName,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const findProductByCategory = `-- name: FindProductByCategory :many
SELECT products.product_id, categories.category_name, products.product_name, products.price, products.created_at FROM products INNER JOIN categories ON products.category_id=categories.category_id WHERE products.category_id=$1
`

type FindProductByCategoryRow struct {
	ProductID    string    `json:"product_id"`
	CategoryName string    `json:"category_name"`
	ProductName  string    `json:"product_name"`
	Price        int64     `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
}

func (q *Queries) FindProductByCategory(ctx context.Context, categoryID int64) ([]FindProductByCategoryRow, error) {
	rows, err := q.db.QueryContext(ctx, findProductByCategory, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindProductByCategoryRow
	for rows.Next() {
		var i FindProductByCategoryRow
		if err := rows.Scan(
			&i.ProductID,
			&i.CategoryName,
			&i.ProductName,
			&i.Price,
			&i.CreatedAt,
		); err != nil {
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
