package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zahrou/ecommerce/db/sqlc"
)

type brandRequest struct {
	Braname  string `json:"braname" binding:"required,alpha"`
	ImageUrl string `json:"imageurl" binding:"required,alpha"`
}

func (server *Server) createBrand(ctx *gin.Context) {
	var req brandRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	braname := db.CreateBrandParams{
		Braname:  req.Braname,
		Imageurl: req.ImageUrl,
	}

	brand, err := server.store.CreateBrand(ctx, braname)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, brand)

}

type listBrandRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getBrands(ctx *gin.Context) {
	var req listBrandRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListBrandsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	brands, err := server.store.ListBrands(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, brands)

}

type brandsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getBrand(ctx *gin.Context) {
	var req brandsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	brand, err := server.store.GetBrand(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, brand)

}

func (server *Server) deleteBrand(ctx *gin.Context) {
	var req brandsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	err := server.store.DeleteBrand(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)

}

type updatebrandsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) updateBrand(ctx *gin.Context) {
	var req updatebrandsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c := ctx.Param("braname")

	arg := db.UpdateBrandParams{
		ID:      req.ID,
		Braname: c,
	}

	brand, err := server.store.UpdateBrand(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, brand)

}
