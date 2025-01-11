package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) listMeals(ctx *gin.Context) {
	// var req listMealRequest
	// if err := ctx.ShouldBindQuery(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }

	meals, err := server.store.ListMeals(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, meals)
}
