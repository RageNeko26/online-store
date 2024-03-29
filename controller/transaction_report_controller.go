package controller

import (
	"context"
	"net/http"

	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/RageNeko26/online-store/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Checkout product
//
//	@Summary Checkout product
//	@Description 	Adding list products in cart and make a transaction report
//	@Tags			Transaction Report
//	@Produce 		json
//	@Param			Authorization header	string true "authorization token customer"
//	@Success		201							{object} CreateTransactionResponse
//	@Failure		500							{object} Response
//	@Router			/checkout [post]
func (controller *Controller) CreateTransaction(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	// Check if authorization is valid
	res, err := utils.DecodeToken(authorization, controller.Secret)

	if err != nil {
		c.Status(http.StatusForbidden)
		return c.JSON(Response{
			Message: "Forbidden",
			Status:  "fail",
		})

	}

	// We need to get list of carts that user have been added
	cart, err := controller.Q.FindCart(context.Background(), res.CustomerID)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return c.JSON(Response{
			Message: "Failed to create transaction because customer is not having cart items",
			Status:  "fail",
		})

	}

	var totalPrice int64

	// Insert into database retrieved items in cart
	for _, items := range cart {
		totalPrice += items.Price
	}

	total, err := controller.Q.CreateTransactionReport(context.Background(), db.CreateTransactionReportParams{
		TransactionID: uuid.NewString(),
		CustomerID:    res.CustomerID,
		TotalPrice:    totalPrice,
	})

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return c.JSON(Response{
			Message: "Failed to create transaction",
			Status:  "fail",
		})

	}

	c.Status(http.StatusCreated)
	return c.JSON(CreateTransactionResponse{
		Products:      cart,
		TransactionID: total.TransactionID,
		TotalPrice:    totalPrice,
		CreatedAt:     total.CreatedAt,
	})
}
