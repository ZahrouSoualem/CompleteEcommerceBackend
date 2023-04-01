package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zahrou/ecommerce/db/sqlc"
)

type CommentRequest struct {
	ReviewID int64  `json:"review_id" binding:"required,min=1"`
	Opinion  string `json:"opinion" binding:"required,min=1"`
}

func (server *Server) createlComment(ctx *gin.Context) {
	var req CommentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCommentParams{
		ReviewID: req.ReviewID,
		Opinion:  req.Opinion,
	}

	Comment, err := server.store.CreateComment(ctx, arg)

	CheckError(err, ctx)

	ctx.JSON(http.StatusOK, Comment)

}

type listCommentsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getComments(ctx *gin.Context) {
	var req listCommentsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCommentsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	Comments, err := server.store.ListComments(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, Comments)

}

type CommentsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getComment(ctx *gin.Context) {
	var req CommentsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	Comment, err := server.store.GetComment(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, Comment)

}

func (server *Server) deleteComment(ctx *gin.Context) {
	var req CommentsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	err := server.store.DeleteComment(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)

}

type updateCommentsRequest struct {
	ID      int64  `uri:"id" binding:"required,min=1"`
	Opinion string `json:"opinion" binding:"required,min=1"`
}

func (server *Server) updateComment(ctx *gin.Context) {

	var req updateCommentsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCommentParams{
		ID:      req.ID,
		Opinion: req.Opinion,
	}

	Comment, err := server.store.UpdateComment(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, Comment)

}
