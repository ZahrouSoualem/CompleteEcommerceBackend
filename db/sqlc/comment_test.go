package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zahrou/ecommerce/util"
)

func CreateRandomComment(t *testing.T) Comment {
	review := CreateRandomReview(t)
	args := CreateCommentParams{
		ReviewID: review.ID,
		Opinion:  util.RnadomName(),
	}

	comment, err := testQueries.CreateComment(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, comment.ReviewID, review.ID)
	require.NotEmpty(t, comment.Opinion)
	return comment

}

func TestComment(t *testing.T) {
	CreateRandomComment(t)
}

func TestGetComment(t *testing.T) {
	comment := CreateRandomComment(t)

	comment2, err := testQueries.GetComment(context.Background(), comment.ID)

	require.NoError(t, err)
	require.NotEmpty(t, comment2)

	require.Equal(t, comment.ID, comment2.ID)
	require.Equal(t, comment.Opinion, comment2.Opinion)
	require.Equal(t, comment.ReviewID, comment2.ReviewID)
}

func TestUpdateComment(t *testing.T) {
	comment := CreateRandomComment(t)

	args := UpdateCommentParams{
		ID:      comment.ID,
		Opinion: util.RandomString(200),
	}

	comment2, err := testQueries.UpdateComment(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, comment2)
	require.Equal(t, args.Opinion, comment2.Opinion)
	require.Equal(t, args.ID, comment2.ID)

}

func TestDeleteComme(t *testing.T) {
	comment := CreateRandomComment(t)

	err := testQueries.DeleteComment(context.Background(), comment.ID)

	require.NoError(t, err)

	comment2, err := testQueries.GetComment(context.Background(), comment.ID)

	require.Error(t, err)
	require.Empty(t, comment2)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}

func TestListComments(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomComment(t)
	}

	args := ListCommentsParams{
		Limit:  5,
		Offset: 5,
	}

	comments, err := testQueries.ListComments(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, comments, 5)

	for _, comment := range comments {
		require.NotEmpty(t, comment)
	}
}
