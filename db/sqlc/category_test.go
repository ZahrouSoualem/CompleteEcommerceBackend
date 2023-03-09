package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zahrou/ecommerce/util"
)

func CreateRandomCategory(t *testing.T) Category {
	arg := util.RnadomName()
	category, err := testQueries.CreateCategory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg, category.Catname)
	require.NotZero(t, category.ID)

	return category
}

func TestGetCategory(t *testing.T) {
	category := CreateRandomCategory(t)

	category2, err := testQueries.GetCategory(context.Background(), category.ID)

	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category2.Catname, category.Catname)
	require.Equal(t, category2.ID, category.ID)
}

func TestCreatCategory(t *testing.T) {
	CreateRandomCategory(t)
}

func TestUpdateCategory(t *testing.T) {
	category1 := CreateRandomCategory(t)
	arg := UpdateCategoryParams{
		ID:      category1.ID,
		Catname: util.RnadomName(),
	}
	category2, err := testQueries.UpdateCategory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.NotZero(t, category2.ID)
	require.Equal(t, arg.Catname, category2.Catname)
}

func TestDeleteCategory(t *testing.T) {
	category1 := CreateRandomCategory(t)
	err := testQueries.DeleteCategories(context.Background(), category1.ID)

	require.NoError(t, err)

	category2, err := testQueries.GetCategory(context.Background(), category1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}

func TestListCategory(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomCategory(t)
	}

	arg := ListCategoriesParams{
		Limit:  5,
		Offset: 5,
	}

	categories, err := testQueries.ListCategories(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, categories, 5)

	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}
