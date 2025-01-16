package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/tijanadmi/food_order/db/sqlc"
)

type createOrderRequest struct {
	Order struct {
		Items []struct {
			MealID      int     `json:"mealid"`
			Name        string  `json:"name"`
			Description string  `json:"description"`
			Price       float64 `json:"price"`
			Image       string  `json:"image"`
			Category    string  `json:"category"`
			CreatedAt   string  `json:"created_at"`
			UpdatedAt   string  `json:"updated_at"`
			Quantity    int     `json:"quantity"`
		} `json:"items" binding:"required"`
		Customer struct {
			Name       string `json:"name"`
			Email      string `json:"email"`
			Street     string `json:"street"`
			PostalCode string `json:"postal-code"`
			City       string `json:"city"`
		} `json:"customer" binding:"required"`
	} `json:"order" binding:"required"`
}

func (server *Server) CreateOrder(ctx *gin.Context) {
	// var req listMealRequest
	// if err := ctx.ShouldBindQuery(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }

	/****/
	var req createOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Mapiranje vrednosti iz req u OrderTxParams
	arg := db.OrderTxParams{
		Items: make([]struct {
			ID       int     `json:"id"`
			Name     string  `json:"name"`
			Price    float64 `json:"price"`
			Quantity int     `json:"quantity"`
		}, len(req.Order.Items)),
		Customer: struct {
			Name       string `json:"name"`
			Email      string `json:"email"`
			Street     string `json:"street"`
			PostalCode string `json:"postal_code"`
			City       string `json:"city"`
		}{
			Name:       req.Order.Customer.Name,
			Email:      req.Order.Customer.Email,
			Street:     req.Order.Customer.Street,
			PostalCode: req.Order.Customer.PostalCode,
			City:       req.Order.Customer.City,
		},
	}

	// Kopiranje svake stavke iz req.Items u arg.Items
	for i, item := range req.Order.Items {
		arg.Items[i] = struct {
			ID       int     `json:"id"`
			Name     string  `json:"name"`
			Price    float64 `json:"price"`
			Quantity int     `json:"quantity"`
		}{
			ID:       item.MealID,
			Name:     item.Name,
			Price:    item.Price,
			Quantity: item.Quantity,
		}
	}
	/****/
	order, err := server.store.OrderTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, order)
}
