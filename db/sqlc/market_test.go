package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zahrou/ecommerce/util"
)

func TestCreateMarket(t *testing.T) {
	CreateRandomMarket(t)
}

func TestGetMarket(t *testing.T) {
	market1 := CreateRandomMarket(t)
	market2, err := testQueries.GetMarket(context.Background(), market1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, market2)

	require.Equal(t, market1.ID, market2.ID)
	require.Equal(t, market1.Email, market2.Email)
	require.Equal(t, market1.Marketname, market2.Marketname)
	require.Equal(t, market1.Password, market2.Password)
}

func CreateRandomMarket(t *testing.T) Market {
	arg := CreateMarketParams{
		Marketname: util.RnadomName(),
		Email:      util.RandomEmail(),
		Password:   util.RnadomName(),
	}

	market, err := testQueries.CreateMarket(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, market)

	require.Equal(t, arg.Marketname, market.Marketname)
	require.Equal(t, arg.Email, market.Email)
	require.Equal(t, arg.Password, market.Password)

	require.NotZero(t, market.ID)

	return market
}

func TestUpdateMarket(t *testing.T) {
	market1 := CreateRandomMarket(t)
	arg := UpdateMarketParams{
		ID:         market1.ID,
		Marketname: util.RnadomName(),
	}
	market2, err := testQueries.UpdateMarket(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, market2)

	require.Equal(t, market1.ID, market2.ID)
	require.Equal(t, market1.Email, market2.Email)
	require.Equal(t, arg.Marketname, market2.Marketname)
	require.Equal(t, market1.Password, market2.Password)
}

func TestDeleteMark(t *testing.T) {
	market1 := CreateRandomMarket(t)

	err := testQueries.DeleteMarket(context.Background(), market1.ID)

	require.NoError(t, err)

	market2, err := testQueries.GetMarket(context.Background(), market1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, market2)
}

func TestListMarket(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomMarket(t)
	}

	arg := ListMarketsParams{
		Limit:  5,
		Offset: 5,
	}

	markets, err := testQueries.ListMarkets(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, markets, 5)

	for _, market := range markets {
		require.NotEmpty(t, market)
	}
}
