// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	CreateBrand(ctx context.Context, arg CreateBrandParams) (Brand, error)
	CreateCategory(ctx context.Context, catname string) (Category, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreateMarket(ctx context.Context, arg CreateMarketParams) (Market, error)
	CreateOrder(ctx context.Context, userID int64) (Order, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateReview(ctx context.Context, arg CreateReviewParams) (Review, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	Createordersproduct(ctx context.Context, arg CreateordersproductParams) (Ordersproduct, error)
	DeleteBrand(ctx context.Context, id int64) error
	DeleteCategories(ctx context.Context, id int64) error
	DeleteComment(ctx context.Context, id int64) error
	DeleteMarket(ctx context.Context, id int64) error
	DeleteOrder(ctx context.Context, id int64) error
	DeleteProducts(ctx context.Context, id int64) error
	DeleteUsers(ctx context.Context, id int64) error
	Deleteordersproduct(ctx context.Context, ordersProductID int64) error
	GetBrand(ctx context.Context, id int64) (Brand, error)
	GetCategory(ctx context.Context, id int64) (Category, error)
	GetComment(ctx context.Context, id int64) (Comment, error)
	GetMarket(ctx context.Context, id int64) (Market, error)
	GetOrders(ctx context.Context, id int64) (Order, error)
	GetProducts(ctx context.Context, id int64) (Product, error)
	GetReview(ctx context.Context, id int64) (Review, error)
	GetUsers(ctx context.Context, id int64) (User, error)
	Getordersproduct(ctx context.Context, ordersProductID int64) (Ordersproduct, error)
	ListBrands(ctx context.Context, arg ListBrandsParams) ([]Brand, error)
	ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error)
	ListComments(ctx context.Context, arg ListCommentsParams) ([]Comment, error)
	ListJoinOrderProducts(ctx context.Context) ([]ListJoinOrderProductsRow, error)
	ListJoinProducts(ctx context.Context, arg ListJoinProductsParams) ([]ListJoinProductsRow, error)
	ListMarkets(ctx context.Context, arg ListMarketsParams) ([]Market, error)
	ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error)
	ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error)
	ListReviews(ctx context.Context, arg ListReviewsParams) ([]Review, error)
	ListUserOrders(ctx context.Context, arg ListUserOrdersParams) ([]Order, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	Listordersproducts(ctx context.Context, arg ListordersproductsParams) ([]ListordersproductsRow, error)
	UpdateBrand(ctx context.Context, arg UpdateBrandParams) (Brand, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error)
	UpdateMarket(ctx context.Context, arg UpdateMarketParams) (Market, error)
	UpdateProductName(ctx context.Context, arg UpdateProductNameParams) (Product, error)
	UpdateReview(ctx context.Context, arg UpdateReviewParams) (Review, error)
	UpdateSoldProduct(ctx context.Context, arg UpdateSoldProductParams) (Product, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
