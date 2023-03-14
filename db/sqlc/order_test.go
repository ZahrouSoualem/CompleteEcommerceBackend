package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomOrder(t *testing.T) Order {
	client := CreateRandomUser(t)

	order, err := testQueries.CreateOrder(context.Background(), client.ID)

	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.NotZero(t, order.ID)
	require.Equal(t, client.ID, order.UserID)
	require.NotZero(t, order.CreatedAt)

	return order
}

func TestCreateOrder(t *testing.T) {
	CreateRandomOrder(t)
}

func TestGetOrder(t *testing.T) {
	order := CreateRandomOrder(t)

	order2, err := testQueries.GetOrders(context.Background(), order.ID)

	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order.ID, order2.ID)
	require.Equal(t, order.UserID, order2.UserID)
	require.NotZero(t, order2.CreatedAt)
}

func TestDeleteOrder(t *testing.T) {
	order := CreateRandomOrder(t)

	err := testQueries.DeleteOrder(context.Background(), order.ID)
	require.NoError(t, err)

	order2, err := testQueries.GetOrders(context.Background(), order.ID)

	require.Error(t, err)
	require.Empty(t, order2)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}

func TestListOrders(t *testing.T) {
	orders := make([]Order, 10)
	for i := 0; i < 10; i++ {
		orders[i] = CreateRandomOrder(t)
	}

	fmt.Println("ID IS :", orders[9].ID)
	fmt.Println("ID IS :", orders[9].UserID)
	/* arg := ListOrdersParams{
		UserID: orders[0].UserID, // using the first order's UserID for the query
		Limit:  5,
		Offset: 5,
	} */

	result, err := testQueries.ListOrders(context.Background(), orders[0].UserID)
	require.NoError(t, err)
	//require.Len(t, result, 5)

	for _, order := range result {
		require.NotEmpty(t, order)
	}
}
