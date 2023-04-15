package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/zahrou/ecommerce/db/sqlc"
	"github.com/zahrou/ecommerce/token"
)

type OrderRequest struct {
	UserID      int64              `json:"userid" binding:"required"`
	ProductList []map[string]int64 `json:"productlist" binding:"required"`
}

type ProductAmount struct {
	ProductID    int64   `json:"productid" `
	ProductName  string  `json:"proname" `
	ProductPrice float64 `json:"proprice" `
	Amount       int64   `json:"proamount" `
	Total        float64 `json:"total"`
}

type OrderReseponse struct {
	ID          int64           `json:"id"`
	UserID      int64           `json:"user_id"`
	CreatedAt   time.Time       `json:"created_at"`
	LastUpdated time.Time       `json:"last_updated"`
	Products    []ProductAmount `json:"soldproduct" `
}

type OrderResponse struct {
	ID          int64                    `json:"id"`
	CreatedAt   time.Time                `json:"created_at"`
	UserID      int64                    `json:"user_id"`
	Username    string                   `json:"username"`
	Email       string                   `json:"email"`
	PhoneNumber int64                    `json:"phone_number"`
	Products    []db.ProductByOrderIDRow `json:"productslist"`
	TotalPrice  float64                  `json:"pricetotal"`
}

type OrderProductResponse struct {
	Productid         int64   `json:"productid"`
	Proname           string  `json:"proname"`
	Price             float64 `json:"price"`
	Quantity          int64   `json:"quantity"`
	PotalProductPrice float64 `json:"total"`
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req OrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, req.UserID)
		return
	}

	arg := req.UserID
	authPayLoad := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUsers(ctx, arg)

	if user.Username != authPayLoad.Username {
		ctx.JSON(http.StatusInternalServerError, arg)
		return
	}

	fmt.Println(req.UserID)

	Order, err := server.store.CreateOrder(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, arg)
		return
	}

	arg2 := req.ProductList

	/* errs := make(chan error)/*/
	results := make(chan []ProductAmount)
	var result db.TxResultParams
	var products []ProductAmount

	go func() {

		for _, value := range arg2 {

			id := value["id"]
			amount := value["amount"]

			product, err := server.store.GetProducts(ctx, id)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, errorResponse(err))
					return
				}
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
			fmt.Println("tha result :", product.Quantity-amount)
			arg3 := db.TransactionTxPrams{
				OrderID:   Order.ID,
				ProductID: id,
				Quantity:  amount,
			}
			result, err = server.store.TransactionTX(ctx, arg3)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
			total := float64(amount) * result.Product.Price
			products = append(products, ProductAmount{
				ProductID:    result.Product.ID,
				ProductName:  result.Product.Proname,
				ProductPrice: result.Product.Price,
				Amount:       amount,
				Total:        total,
			})

		}
		results <- products

	}()

	response := OrderReseponse{
		ID:          Order.ID,
		UserID:      Order.UserID,
		CreatedAt:   Order.CreatedAt,
		LastUpdated: Order.LastUpdated,
		Products:    <-results,
	}

	ctx.JSON(http.StatusOK, response)

}

func (server *Server) getDetailedOrders(ctx *gin.Context) {

	Orders, err := server.store.ListJoinOrderProducts(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var Response []OrderResponse
	for _, value := range Orders {
		OrderProduct, err := server.store.ProductByOrderID(ctx, value.Orderid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		var sum = 0.0
		for _, value3 := range OrderProduct {
			sum = sum + float64(value3.Price)*float64(value3.Quantity)
		}
		Response = append(Response, OrderResponse{
			ID:          value.Orderid,
			CreatedAt:   value.CreatedAt,
			UserID:      value.Userid,
			Username:    value.Username,
			Email:       value.Email,
			PhoneNumber: value.PhoneNumber,
			Products:    OrderProduct,
			TotalPrice:  sum,
		})

	}

	ctx.JSON(http.StatusOK, Response)

}

type OrdersRequest struct {
	UserID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetOrdersByUserID(ctx *gin.Context) {

	var req OrdersRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.UserID
	authPayLoad := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUsers(ctx, arg)

	if user.Username != authPayLoad.Username {
		ctx.JSON(http.StatusInternalServerError, arg)
		return
	}

	Orders, err := server.store.ListOrderByUserID(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var Response []OrderResponse
	for _, value := range Orders {
		OrderProduct, err := server.store.ProductByOrderID(ctx, value.Orderid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		var sum = 0.0
		for _, value3 := range OrderProduct {
			sum = sum + float64(value3.Price)*float64(value3.Quantity)
		}
		Response = append(Response, OrderResponse{
			ID:          value.Orderid,
			CreatedAt:   value.CreatedAt,
			UserID:      value.Userid,
			Username:    authPayLoad.Username,
			Email:       value.Email,
			PhoneNumber: value.PhoneNumber,
			Products:    OrderProduct,
			TotalPrice:  sum,
		})

	}

	ctx.JSON(http.StatusOK, Response)

}
