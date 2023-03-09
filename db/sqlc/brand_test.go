package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zahrou/ecommerce/util"
)

func CreateRandomBrand(t *testing.T) Brand {

	arg := CreateBrandParams{
		Braname:  util.RnadomName(),
		Imageurl: util.RnadomName(),
	}
	brand, err := testQueries.CreateBrand(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, brand)

	require.Equal(t, arg.Braname, brand.Braname)
	require.Equal(t, arg.Imageurl, brand.Imageurl)

	require.NotZero(t, brand.ID)

	return brand
}

func TestCreateBrand(t *testing.T) {
	CreateRandomBrand(t)
}

func TestUpdateBrand(t *testing.T) {
	brand := CreateRandomBrand(t)
	arg := UpdateBrandParams{
		ID:      brand.ID,
		Braname: util.RnadomName(),
	}
	brand2, err := testQueries.UpdateBrand(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, brand2)

	require.Equal(t, brand.ID, brand2.ID)
	require.Equal(t, arg.Braname, brand2.Braname)

	require.NotZero(t, brand2)
}

func TestDeleteBrand(t *testing.T) {
	brand := CreateRandomBrand(t)
	err := testQueries.DeleteBrand(context.Background(), brand.ID)

	require.NoError(t, err)
	brand2, err := testQueries.GetBrand(context.Background(), brand.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, brand2)
}

func TestGetBrand(t *testing.T) {
	brand := CreateRandomBrand(t)

	brand2, err := testQueries.GetBrand(context.Background(), brand.ID)

	require.NoError(t, err)
	require.NotEmpty(t, brand2)

	require.Equal(t, brand.Braname, brand2.Braname)
	require.Equal(t, brand.Imageurl, brand2.Imageurl)
}

func TestListBrands(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomBrand(t)
	}

	args := ListBrandsParams{
		Limit:  5,
		Offset: 5,
	}

	brands, err := testQueries.ListBrands(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, brands, 5)

	for _, brand := range brands {

		require.NotEmpty(t, brand)
	}

}
