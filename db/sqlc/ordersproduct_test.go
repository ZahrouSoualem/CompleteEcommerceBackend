package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zahrou/ecommerce/util"
)

func CreateRandomOrderProduct(t *testing.T, orderID int64) Ordersproduct {

	Product := CreateRandomProduct(t)
	args := CreateordersproductParams{
		OrdersID:  orderID,
		ProductID: Product.ID,
		Quantity:  util.RandomInteger(9, 20),
	}

	orderproduct, err := testQueries.Createordersproduct(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, orderproduct)

	require.Equal(t, args.OrdersID, orderproduct.OrdersID)
	require.Equal(t, args.ProductID, orderproduct.ProductID)

	return orderproduct

}

func TestCreateOrderProduct(t *testing.T) {
	CreateRandomOrderProduct(t, CreateRandomOrder(t).ID)
}

func TestCreateListOrderProduct(t *testing.T) {
	//var order Order
	order := CreateRandomOrder(t)

	for i := 0; i < 10; i++ {
		CreateRandomOrderProduct(t, order.ID)
	}

}

func TestGetOProduct(t *testing.T) {
	order := CreateRandomOrder(t)
	orderproduct := CreateRandomOrderProduct(t, order.ID)

	order2, err := testQueries.Getordersproduct(context.Background(), orderproduct.OrdersProductID)

	require.NoError(t, err)
	require.NotEmpty(t, order2)
}

func TestListOrderProductByID(t *testing.T) {
	order := CreateRandomOrder(t)
	for i := 0; i < 10; i++ {
		CreateRandomOrderProduct(t, order.ID)
	}

	args := ListordersproductsParams{
		OrdersID: order.ID,
		Limit:    5,
		Offset:   5,
	}

	orderproduct, err := testQueries.Listordersproducts(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, orderproduct, 5)

	for _, oprodcut := range orderproduct {
		require.NotEmpty(t, oprodcut)
	}
}
