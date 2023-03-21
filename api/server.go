package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/zahrou/ecommerce/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {

	server := &Server{store: store}
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

	server.router = router

	return server

}

// Starts runs the HTTP server on specific address

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Starts runs the HTTP server on specific address

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
