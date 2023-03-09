// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: comment.sql

package db

import (
	"context"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comment (
  review_id,
  opinion
) VALUES (
  $1, $2
)
RETURNING id, review_id, opinion
`

type CreateCommentParams struct {
	ReviewID int64  `json:"review_id"`
	Opinion  string `json:"opinion"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment, arg.ReviewID, arg.Opinion)
	var i Comment
	err := row.Scan(&i.ID, &i.ReviewID, &i.Opinion)
	return i, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM Comment
WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteComment, id)
	return err
}

const getComment = `-- name: GetComment :one
SELECT id, review_id, opinion FROM comment
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetComment(ctx context.Context, id int64) (Comment, error) {
	row := q.db.QueryRowContext(ctx, getComment, id)
	var i Comment
	err := row.Scan(&i.ID, &i.ReviewID, &i.Opinion)
	return i, err
}

const listComments = `-- name: ListComments :many
SELECT id, review_id, opinion FROM comment
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListCommentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListComments(ctx context.Context, arg ListCommentsParams) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, listComments, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(&i.ID, &i.ReviewID, &i.Opinion); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateComment = `-- name: UpdateComment :one
UPDATE comment SET opinion = $2
WHERE id = $1
RETURNING id, review_id, opinion
`

type UpdateCommentParams struct {
	ID      int64  `json:"id"`
	Opinion string `json:"opinion"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, updateComment, arg.ID, arg.Opinion)
	var i Comment
	err := row.Scan(&i.ID, &i.ReviewID, &i.Opinion)
	return i, err
}
