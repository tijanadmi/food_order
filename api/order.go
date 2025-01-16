package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/tijanadmi/food_order/db/sqlc"
)

type createOrderRequest struct {
	Items []struct {
		ID       int     `json:"id"`
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity"`
	} `json:"items" binding:"required"`
	Customer struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		Street     string `json:"street"`
		PostalCode string `json:"postal_code"`
		City       string `json:"city"`
	} `json:"customer" binding:"required"`
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
		}, len(req.Items)),
		Customer: struct {
			Name       string `json:"name"`
			Email      string `json:"email"`
			Street     string `json:"street"`
			PostalCode string `json:"postal_code"`
			City       string `json:"city"`
		}{
			Name:       req.Customer.Name,
			Email:      req.Customer.Email,
			Street:     req.Customer.Street,
			PostalCode: req.Customer.PostalCode,
			City:       req.Customer.City,
		},
	}

	// Kopiranje svake stavke iz req.Items u arg.Items
	for i, item := range req.Items {
		arg.Items[i] = struct {
			ID       int     `json:"id"`
			Name     string  `json:"name"`
			Price    float64 `json:"price"`
			Quantity int     `json:"quantity"`
		}{
			ID:       item.ID,
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
