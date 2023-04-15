package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/zahrou/ecommerce/db/sqlc"
	"github.com/zahrou/ecommerce/token"
	"github.com/zahrou/ecommerce/util"
)

type Server struct {
	config util.Config
	// So inside the server you have to put the interface that hold the queries
	store     db.Store
	toenMaker token.Maker
	// also don't forget the engine the will manage the routing
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{store: store, toenMaker: tokenMaker}
	router := gin.Default()

	//category
	router.POST("/category", server.createCategory)
	router.GET("/categories", server.getCategories)
	router.GET("/category/:id", server.getCategory)
	router.DELETE("/category/:id", server.deleteCategory)
	router.PUT("/category/:id/:catname", server.updateCategory)

	//brand
	router.POST("/brand", server.createBrand)
	router.GET("/brands", server.getBrands)
	router.GET("/brand/:id", server.getBrand)
	router.DELETE("/brand/:id", server.deleteBrand)
	router.PUT("/brand/:id/:braname", server.updateBrand)

	//Market
	router.POST("/market", server.createmarket)
	router.GET("/markets", server.getmarkets)
	router.GET("/market/:id", server.getmarket)
	router.DELETE("/market/:id", server.deletemarket)
	router.PUT("/market/:id/:marketname", server.updatemarket)

	//product
	router.POST("/product", server.createproduct)
	router.GET("/products", server.getproducts)
	router.GET("/product/:id", server.getproduct)
	router.DELETE("/product/:id", server.deleteproduct)
	router.PUT("/product/:id/:proname", server.updateproduct)

	//review
	router.POST("/review", server.createlReview)
	router.GET("/reviews", server.getReviews)
	router.GET("/review/:id", server.getReview)
	//router.DELETE("/review/:id", server.deletereview)
	router.PUT("/review/:id/:rating", server.updateReview)

	//Comment
	router.POST("/comment", server.createlComment)
	router.GET("/comments", server.getComments)
	router.GET("/comment/:id", server.getComment)
	router.DELETE("/comment/:id", server.deleteComment)
	router.PUT("/comment", server.updateComment)

	//user
	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)
	router.GET("/users", server.getUsers)

	authRouter := router.Group("/").Use(authMiddleWare(tokenMaker))

	authRouter.GET("/user/:id", server.getUser)
	authRouter.DELETE("/user/:id", server.deleteUser)
	authRouter.PUT("/user/:id/:username", server.updateUser)

	authRouter.POST("/order", server.createOrder)
	authRouter.GET("/orderslist", server.getDetailedOrders)
	authRouter.GET("/orderslist/:id", server.GetOrdersByUserID)

	server.router = router

	return server, nil

}

// Starts runs the HTTP server on specific address

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Starts runs the HTTP server on specific address

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func CheckError(err error, ctx *gin.Context) {
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation", "users_phone_number_key":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
}
