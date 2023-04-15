package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TransactionTX(ctx context.Context, arg TransactionTxPrams) (TxResultParams, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db), // to get sql queries
		db:      db,      // to get Tx
	}
}


func (store *SQLStore) execTX(ctx context.Context, fn func(*Queries) error) error {

	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx rollback error: %v, original error: %v", rbErr, err)
		}
		return err
	}

	return tx.Commit()

}

type TransactionTxPrams struct {
	OrderID   int64
	ProductID int64
	Quantity  int64
}

type TxResultParams struct {
	OrderProduct Ordersproduct
	Product      Product
}

func (store *SQLStore) TransactionTX(ctx context.Context, arg TransactionTxPrams) (TxResultParams, error) {

	var result TxResultParams
	err := store.execTX(ctx, func(q *Queries) error {
		var err error

		result.OrderProduct, err = q.Createordersproduct(ctx, CreateordersproductParams{
			OrdersID:  arg.OrderID,
			ProductID: arg.ProductID,
			Quantity:  arg.Quantity,
		})

		if err != nil {
			return err
		}

		//Update the Product quantity

		result.Product, err = q.UpdateSoldProduct(ctx, UpdateSoldProductParams{
			Amount: -arg.Quantity,
			ID:     arg.ProductID,
		})

		if err != nil {
			fmt.Printf("%v", err)
			return err
		}

		return nil

	})

	return result, err

}
