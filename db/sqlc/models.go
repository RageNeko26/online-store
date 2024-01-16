// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"time"
)

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Customer struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Product struct {
	ID         string    `json:"id"`
	CategoryID string    `json:"category_id"`
	Name       string    `json:"name"`
	Price      int64     `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
}

type ShoppingCart struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	ProductID  string    `json:"product_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type TransactionReport struct {
	ID        string         `json:"id"`
	CartID    sql.NullString `json:"cart_id"`
	CreatedAt time.Time      `json:"created_at"`
}
