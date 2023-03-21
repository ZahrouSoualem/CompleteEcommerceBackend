package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zahrou/ecommerce/db/sqlc"
)

type marketRequest struct {
	Marketname string `json:"marketname" binding:"required,alpha"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

func (server *Server) createmarket(ctx *gin.Context) {
	var req marketRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMarketParams{
		Marketname: req.Marketname,
		Email:      req.Email,
		Password:   req.Password,
	}

	market, err := server.store.CreateMarket(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, market)

}

type listmarketRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getmarkets(ctx *gin.Context) {
	var req listmarketRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListMarketsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	markets, err := server.store.ListMarkets(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, markets)

}

type marketsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getmarket(ctx *gin.Context) {
	var req marketsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	market, err := server.store.GetMarket(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, market)

}

func (server *Server) deletemarket(ctx *gin.Context) {
	var req marketsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	err := server.store.DeleteMarket(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)

}

type updatemarketsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) updatemarket(ctx *gin.Context) {
	var req updatemarketsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c := ctx.Param("marketname")

	arg := db.UpdateMarketParams{
		ID:         req.ID,
		Marketname: c,
	}

	market, err := server.store.UpdateMarket(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, market)

}
