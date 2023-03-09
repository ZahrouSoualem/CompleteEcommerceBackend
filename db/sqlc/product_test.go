package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zahrou/ecommerce/util"
)

func CreateRandomProduct(t *testing.T) Product {
	category := CreateRandomCategory(t)
	brand := CreateRandomBrand(t)
	market := CreateRandomMarket(t)
	args := CreateProductParams{
		CategoryID:  category.ID,
		BrandID:     brand.ID,
		MarketID:    market.ID,
		Proname:     util.RnadomName(),
		Description: util.RandomString(68),
		Imageurl:    util.RnadomName(),
		Price:       float64(util.RandomInteger(50, 100)),
		Quantity:    util.RandomInteger(50, 100),
	}

	product, err := testQueries.CreateProduct(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, args.BrandID, product.BrandID)
	require.Equal(t, args.CategoryID, product.CategoryID)
	require.Equal(t, args.MarketID, product.MarketID)
	require.Equal(t, args.Imageurl, product.Imageurl)
	require.Equal(t, args.Price, product.Price)
	require.Equal(t, args.Proname, product.Proname)
	require.Equal(t, args.Quantity, product.Quantity)

	require.NotZero(t, product.ID)

	return product
}

func TestGetProduct(t *testing.T) {
	product := CreateRandomProduct(t)

	product2, err := testQueries.GetProducts(context.Background(), product.ID)

	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product2.Proname, product.Proname)
	require.Equal(t, product2.ID, product.ID)

}

func TestCreateProduct(t *testing.T) {
	CreateRandomProduct(t)
}

func TestDeleteProduct(t *testing.T) {
	product := CreateRandomProduct(t)

	err := testQueries.DeleteProducts(context.Background(), product.ID)

	require.NoError(t, err)

	product2, err := testQueries.GetProducts(context.Background(), product.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, product2)
}

func TestListProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomProduct(t)
	}

	args := ListProductsParams{
		Limit:  5,
		Offset: 5,
	}

	products, err := testQueries.ListProducts(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, products, 5)

	for _, product := range products {
		require.NotEmpty(t, product)
	}
}

/* func TestUpdateProduct(t *testing.T) {
	product := CreateRandomProduct(t)

	args := UpdateProductParams{
		ID: product.ID,
		Proname: util.RandomString(12),
	}

	product2,err := testQueries.UpdateProduct(context.Background(),args)

	require.NoError(t,err)
	require.Equal(t,args.Proname,product2.Proname)

	require.


} */
