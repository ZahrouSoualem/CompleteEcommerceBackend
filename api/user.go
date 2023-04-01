package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/zahrou/ecommerce/db/sqlc"
	"github.com/zahrou/ecommerce/util"
)

type userRequest struct {
	Username    string `json:"username" binding:"required,alpha"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	ZipCode     int64  `json:"zip_code" binding:"required,numeric"`
	PhoneNumber int64  `json:"phone_number" binding:"required,numeric"`
}

type userResponse struct {
	ID          int64     `json:"id" `
	Username    string    `json:"username" `
	Email       string    `json:"email" `
	Address     string    `json:"address"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	ZipCode     int64     `json:"zip_code"`
	PhoneNumber int64     `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req userRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedpassword, err := util.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user := db.CreateUserParams{
		Username:    req.Username,
		Email:       req.Email,
		Password:    hashedpassword,
		Address:     req.Address,
		City:        req.City,
		State:       req.State,
		Country:     req.Country,
		ZipCode:     req.ZipCode,
		PhoneNumber: req.PhoneNumber,
	}

	User, err := server.store.CreateUser(ctx, user)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation", "users_phone_number_key":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
	}

	returneduser := userResponse{
		ID:          User.ID,
		Username:    User.Username,
		Email:       User.Email,
		Address:     User.Address,
		City:        User.City,
		State:       User.State,
		Country:     User.Country,
		ZipCode:     User.ZipCode,
		PhoneNumber: User.PhoneNumber,
		CreatedAt:   User.CreatedAt,
	}

	ctx.JSON(http.StatusOK, returneduser)

}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getUsers(ctx *gin.Context) {
	var req listUserRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	Users, err := server.store.ListUsers(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, Users)

}

type UsersRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req UsersRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	User, err := server.store.GetUsers(ctx, arg)

	if err != nil {

		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation", "users_phone_number_key":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		/* for err != nil && ok && pqErr.Code.Name() == "unique_violation" && strings.Contains(pqErr.Constraint, "users_phone_number_key") {

		} */

	}

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	returneduser := userResponse{
		ID:          User.ID,
		Username:    User.Username,
		Email:       User.Email,
		Address:     User.Address,
		City:        User.City,
		State:       User.State,
		Country:     User.Country,
		ZipCode:     User.ZipCode,
		PhoneNumber: User.PhoneNumber,
		CreatedAt:   User.CreatedAt,
	}

	ctx.JSON(http.StatusOK, returneduser)

}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req UsersRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.ID

	err := server.store.DeleteUsers(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)

}

type updateUsersRequest struct {
	ID       int64  `uri:"id" binding:"required,min=1"`
	Username string `uri:"username" binding:"required,alpha"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUsersRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//c := ctx.Param("braname")

	arg := db.UpdateUserParams{
		ID:       req.ID,
		Username: req.Username,
	}

	User, err := server.store.UpdateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, User)

}

type loginuserRequest struct {
	Username string `json:"username" binding:"required,alpha"`
	Password string `json:"password" binding:"required"`
}

type loginuserResponse struct {
	Token string       `json:"access_token" `
	User  userResponse `json:"user" `
}

func (server *Server) loginUser(ctx *gin.Context) {

	var req loginuserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	User, err := server.store.GetUsersByID(ctx, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, User.Password)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	accessToen, _, err := server.toenMaker.CreateToken(req.Username, server.config.AccessDuration)

	resq := loginuserResponse{
		Token: accessToen,
		User: userResponse{
			ID:          User.ID,
			Username:    User.Username,
			Email:       User.Email,
			Address:     User.Address,
			City:        User.City,
			State:       User.State,
			Country:     User.Country,
			ZipCode:     User.ZipCode,
			PhoneNumber: User.PhoneNumber,
			CreatedAt:   User.CreatedAt,
		},
	}
	ctx.JSON(http.StatusOK, resq)

}
