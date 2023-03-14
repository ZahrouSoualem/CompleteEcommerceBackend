package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zahrou/ecommerce/util"
)

func CreateRandomReview(t *testing.T) Review {
	client := CreateRandomUser(t)
	product := CreateRandomProduct(t)
	args := CreateReviewParams{
		UserID:    client.ID,
		ProductID: product.ID,
		Rating:    util.RandomInteger(0, 5),
	}

	review, err := testQueries.CreateReview(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, review)

	require.Equal(t, review.UserID, client.ID)
	require.Equal(t, review.ProductID, product.ID)
	require.NotEmpty(t, review.ID)
	//require.NotEmpty(t, review.Rating)

	return review
}

func TestCreateReview(t *testing.T) {
	CreateRandomReview(t)
}
func TestGetReview(t *testing.T) {
	review := CreateRandomReview(t)

	review2, err := testQueries.GetReview(context.Background(), review.ID)

	require.NoError(t, err)
	require.NotEmpty(t, review2)

	require.Equal(t, review.ID, review2.ID)
	require.Equal(t, review.ProductID, review2.ProductID)
	require.Equal(t, review.Rating, review2.Rating)
	require.Equal(t, review.UserID, review2.UserID)
}
func TestUpdateReview(t *testing.T) {
	review := CreateRandomReview(t)
	args := UpdateReviewParams{
		ID:     review.ID,
		Rating: util.RandomInteger(0, 5),
	}
	review2, err := testQueries.UpdateReview(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, review2)

	require.Equal(t, args.ID, review2.ID)
	require.Equal(t, args.Rating, review2.Rating)
	require.NotEmpty(t, review2.ID)

}
