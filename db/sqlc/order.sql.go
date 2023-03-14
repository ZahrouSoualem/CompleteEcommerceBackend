// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: order.sql

package db

import (
	"context"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
  user_id
) VALUES (
  $1
)
RETURNING id, user_id, created_at, last_updated
`

func (q *Queries) CreateOrder(ctx context.Context, userID int64) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder, userID)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, id)
	return err
}

const getOrders = `-- name: GetOrders :one
SELECT id, user_id, created_at, last_updated FROM orders
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOrders(ctx context.Context, id int64) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrders, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT id, user_id, created_at, last_updated FROM orders
WHERE user_id=$1
`

func (q *Queries) ListOrders(ctx context.Context, userID int64) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrders, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CreatedAt,
			&i.LastUpdated,
		); err != nil {
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
