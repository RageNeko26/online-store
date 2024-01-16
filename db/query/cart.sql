-- name: CreateCart :one
INSERT into shopping_carts (
    cart_id,
    customer_id, 
    product_id
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: FindCart :many
SELECT shopping_carts.cart_id, customers.customer_name, products.product_name FROM shopping_carts
    INNER JOIN customers ON shopping_carts.customer_id=customers.customer_id
    INNER JOIN products on shopping_carts.product_id=products.product_id
    WHERE customers.customer_name=$1;