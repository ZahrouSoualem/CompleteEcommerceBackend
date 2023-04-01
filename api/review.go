package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zahrou/ecommerce/db/sqlc"
)

type reviewRequest struct {
	UserID    int64 `json:"user_id" binding:"required,min=1"`
	ProductID int64 `json:"product_id" binding:"required,min=1"`
	Rating    int64 `json:"rating" binding:"required,min=0,max=5"`
}

func (server *Server) createlReview(ctx *gin.Context) {
	var req reviewRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateReviewParams{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Rating:    req.Rating,
	}

	review, err := server.store.CreateReview(ctx, arg)

	CheckError(err, ctx)

	ctx.JSON(http.StatusOK, review)

}

type listReviewsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getReviews(ctx *gin.Context) {
	var req listReviewsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListReviewsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	Reviews, err := server.store.ListReviews(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Println(Reviews)
	ctx.JSON(http.StatusOK, Reviews)

}

type ReviewsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getReview(ctx *gin.Context) {
	var req ReviewsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	review, err := server.store.GetReview(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, review)

}

/*func (server *Server) deleteReview(ctx *gin.Context) {
	var req ReviewsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	err := server.store.(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)

}*/

type updateReviewsRequest struct {
	ID     int64 `uri:"id" binding:"required,min=1"`
	Rating int64 `uri:"catname" binding:"required,min=0,max=5"`
}

func (server *Server) updateReview(ctx *gin.Context) {

	var req updateReviewsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateReviewParams{
		ID:     req.ID,
		Rating: req.Rating,
	}

	review, err := server.store.UpdateReview(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, review)

}
