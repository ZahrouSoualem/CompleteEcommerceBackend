package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zahrou/ecommerce/db/sqlc"
)

type productRequest struct {
	CategoryID  int64   `json:"category_id" binding:"required,min=1" `
	BrandID     int64   `json:"brand_id" binding:"required,min=1"`
	MarketID    int64   `json:"market_id" binding:"required,min=1"`
	Proname     string  `json:"proname" binding:"required,alpha"`
	Description string  `json:"description" binding:"required"`
	Imageurl    string  `json:"imageurl" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Quantity    int64   `json:"quantity" binding:"required,min=1"`
}

func (server *Server) createproduct(ctx *gin.Context) {
	var req productRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProductParams{
		CategoryID:  req.CategoryID,
		BrandID:     req.BrandID,
		MarketID:    req.MarketID,
		Proname:     req.Proname,
		Description: req.Description,
		Imageurl:    req.Imageurl,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	product, err := server.store.CreateProduct(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)

}

type listproductRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getproducts(ctx *gin.Context) {
	var req listproductRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListJoinProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	products, err := server.store.ListJoinProducts(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, products)

}

type productsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getproduct(ctx *gin.Context) {
	var req productsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	product, err := server.store.GetProducts(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)

}

func (server *Server) deleteproduct(ctx *gin.Context) {
	var req productsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	err := server.store.DeleteProducts(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)

}

type updateproductsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) updateproduct(ctx *gin.Context) {
	var req updateproductsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c := ctx.Param("proname")

	arg := db.UpdateProductNameParams{
		ID:      req.ID,
		Proname: c,
	}

	product, err := server.store.UpdateProductName(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)

}
