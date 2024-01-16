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
    id,
    category_id,
    name,
    price
) VALUES (
    $1, $2, $3, $4
) RETURNING id, category_id, name, price, created_at
`

type CreateProductParams struct {
	ID         string `json:"id"`
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
	Price      int64  `json:"price"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.ID,
		arg.CategoryID,
		arg.Name,
		arg.Price,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Name,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const findProductByCategory = `-- name: FindProductByCategory :many
SELECT products.id, categories.name, products.name, products.price, products.created_at FROM products INNER JOIN categories ON products.category_id=categories.id
`

type FindProductByCategoryRow struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Name_2    string    `json:"name_2"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) FindProductByCategory(ctx context.Context) ([]FindProductByCategoryRow, error) {
	rows, err := q.db.QueryContext(ctx, findProductByCategory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindProductByCategoryRow
	for rows.Next() {
		var i FindProductByCategoryRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Name_2,
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