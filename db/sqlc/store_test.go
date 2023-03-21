package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zahrou/ecommerce/util"
)

func TestTransaction(t *testing.T) {
	store := NewStore(testDB)

	var products []Product
	var orders []Order

	for i := 0; i < 5; i++ {
		products = append(products, CreateRandomProduct(t))
		fmt.Println("quantity of ", products[i].ID, " is = ", products[i].Quantity)
	}

	for i := 0; i < 3; i++ {
		orders = append(orders, CreateRandomOrder(t))
		fmt.Println("order of ", i, " is = ", orders[i].ID)

	}

	// concurrent transaction
	n := 3

	errs := make(chan error)
	results := make(chan TxResultParams)
	var result TxResultParams
	var err error
	for i := 0; i < n; i++ {
		go func(j int) {
			fmt.Println("we are inside ", j)
			for k := 0; k < 5; k++ {
				q := util.RandomInteger(21, 40)
				fmt.Println("tha quantity :", q)
				fmt.Println("tha result :", products[k].Quantity-q)
				result, err = store.TransactionTX(context.Background(), TransactionTxPrams{
					orderID:   orders[j].ID,
					productID: products[k].ID,
					quantity:  q,
				})
			}

			fmt.Println("we are inside the middle of  ", j)

			errs <- err
			results <- result
			fmt.Println("we are inside the end of  ", j)

		}(i)
	}

	var order3 Order
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		orderproduct := result.OrderProduct

		require.NotEmpty(t, orderproduct)
		_, err = testQueries.Getordersproduct(context.Background(), orderproduct.OrdersProductID)
		require.NoError(t, err)

		require.NotEmpty(t, orderproduct)
		_, err = testQueries.GetProducts(context.Background(), orderproduct.ProductID)
		require.NoError(t, err)

		require.NotEmpty(t, orderproduct)
		order3, err = testQueries.GetOrders(context.Background(), orderproduct.OrdersID)
		require.NoError(t, err)
		require.NotEmpty(t, order3)

		_, err = testQueries.GetUsers(context.Background(), order3.UserID)
		require.NoError(t, err)
	}

}

func TestTransaction2(t *testing.T) {
	store := NewStore(testDB)

	var products []Product
	var orders []Order

	for i := 0; i < 1; i++ {
		products = append(products, CreateRandomProduct(t))
		fmt.Println("quantity of ", products[i].ID, " is = ", products[i].Quantity, " - ")
	}

	for i := 0; i < 3; i++ {
		orders = append(orders, CreateRandomOrder(t))
		fmt.Println("order of ", i, " is = ", orders[i].ID)

	}

	// concurrent transaction
	n := 3

	errs := make(chan error)
	results := make(chan TxResultParams)
	var result TxResultParams
	var err error
	for i := 0; i < n; i++ {
		go func(j int) {
			fmt.Println("we are inside ", j)
			for k := 0; k < 1; k++ {
				q := util.RandomInteger(21, 40)
				fmt.Println("tha quantity :", q)
				fmt.Println("tha result :", products[k].Quantity-q)
				result, err = store.TransactionTX(context.Background(), TransactionTxPrams{
					orderID:   orders[j].ID,
					productID: products[k].ID,
					quantity:  q,
				})
			}

			fmt.Println("we are inside the middle of  ", j)

			errs <- err
			results <- result
			fmt.Println("we are inside the end of  ", j)

		}(i)
	}

	var order3 Order
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		orderproduct := result.OrderProduct

		require.NotEmpty(t, orderproduct)
		_, err = testQueries.Getordersproduct(context.Background(), orderproduct.OrdersProductID)
		require.NoError(t, err)

		require.NotEmpty(t, orderproduct)
		_, err = testQueries.GetProducts(context.Background(), orderproduct.ProductID)
		require.NoError(t, err)

		require.NotEmpty(t, orderproduct)
		order3, err = testQueries.GetOrders(context.Background(), orderproduct.OrdersID)
		require.NoError(t, err)
		require.NotEmpty(t, order3)

		_, err = testQueries.GetUsers(context.Background(), order3.UserID)
		require.NoError(t, err)
	}

}
