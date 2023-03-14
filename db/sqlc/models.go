// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"
)

type Brand struct {
	ID       int64  `json:"id"`
	Braname  string `json:"braname"`
	Imageurl string `json:"imageurl"`
}

type Category struct {
	ID      int64  `json:"id"`
	Catname string `json:"catname"`
}

type Comment struct {
	ID       int64  `json:"id"`
	ReviewID int64  `json:"review_id"`
	Opinion  string `json:"opinion"`
}

type Market struct {
	ID         int64  `json:"id"`
	Marketname string `json:"marketname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type Order struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdated time.Time `json:"last_updated"`
}

type Ordersproduct struct {
	OrdersProductID int64         `json:"orders_product_id"`
	OrdersID        int64         `json:"orders_id"`
	ProductID       int64         `json:"product_id"`
	Quantity        sql.NullInt64 `json:"quantity"`
}

type Product struct {
	ID          int64   `json:"id"`
	CategoryID  int64   `json:"category_id"`
	BrandID     int64   `json:"brand_id"`
	MarketID    int64   `json:"market_id"`
	Proname     string  `json:"proname"`
	Description string  `json:"description"`
	Imageurl    string  `json:"imageurl"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
}

type Review struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"product_id"`
	UserID    int64 `json:"user_id"`
	Rating    int64 `json:"rating"`
}

type User struct {
	ID          int64          `json:"id"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	Address     sql.NullString `json:"address"`
	City        sql.NullString `json:"city"`
	State       sql.NullString `json:"state"`
	Country     sql.NullString `json:"country"`
	ZipCode     int64          `json:"zip_code"`
	PhoneNumber int64          `json:"phone_number"`
	CreatedAt   sql.NullTime   `json:"created_at"`
}
