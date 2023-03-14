// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: market.sql

package db

import (
	"context"
)

const createMarket = `-- name: CreateMarket :one
INSERT INTO market (
  marketname,
   email,
   password
   
) VALUES (
  $1, $2, $3
)
RETURNING id, marketname, email, password
`

type CreateMarketParams struct {
	Marketname string `json:"marketname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func (q *Queries) CreateMarket(ctx context.Context, arg CreateMarketParams) (Market, error) {
	row := q.db.QueryRowContext(ctx, createMarket, arg.Marketname, arg.Email, arg.Password)
	var i Market
	err := row.Scan(
		&i.ID,
		&i.Marketname,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const deleteMarket = `-- name: DeleteMarket :exec
DELETE FROM market
WHERE id = $1
`

func (q *Queries) DeleteMarket(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMarket, id)
	return err
}

const getMarket = `-- name: GetMarket :one
SELECT id, marketname, email, password FROM market
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMarket(ctx context.Context, id int64) (Market, error) {
	row := q.db.QueryRowContext(ctx, getMarket, id)
	var i Market
	err := row.Scan(
		&i.ID,
		&i.Marketname,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const listMarkets = `-- name: ListMarkets :many
SELECT id, marketname, email, password FROM market
ORDER BY marketname
LIMIT $1
OFFSET $2
`

type ListMarketsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListMarkets(ctx context.Context, arg ListMarketsParams) ([]Market, error) {
	rows, err := q.db.QueryContext(ctx, listMarkets, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Market
	for rows.Next() {
		var i Market
		if err := rows.Scan(
			&i.ID,
			&i.Marketname,
			&i.Email,
			&i.Password,
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

const updateMarket = `-- name: UpdateMarket :one
UPDATE market SET marketname = $2
WHERE id = $1
RETURNING id, marketname, email, password
`

type UpdateMarketParams struct {
	ID         int64  `json:"id"`
	Marketname string `json:"marketname"`
}

func (q *Queries) UpdateMarket(ctx context.Context, arg UpdateMarketParams) (Market, error) {
	row := q.db.QueryRowContext(ctx, updateMarket, arg.ID, arg.Marketname)
	var i Market
	err := row.Scan(
		&i.ID,
		&i.Marketname,
		&i.Email,
		&i.Password,
	)
	return i, err
}