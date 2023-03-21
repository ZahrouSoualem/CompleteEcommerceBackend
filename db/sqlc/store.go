package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {

	return &Store{
		Queries: New(db), // to get sql queries
		db:      db,      // to get Tx
	}
}

func (store *Store) execTX(ctx context.Context, fn func(*Queries) error) error {

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
	orderID   int64
	productID int64
	quantity  int64
}

type TxResultParams struct {
	OrderProduct Ordersproduct
	product      Product
}

func (store *Store) TransactionTX(ctx context.Context, arg TransactionTxPrams) (TxResultParams, error) {

	var result TxResultParams
	err := store.execTX(ctx, func(q *Queries) error {
		var err error

		result.OrderProduct, err = q.Createordersproduct(ctx, CreateordersproductParams{
			OrdersID:  arg.orderID,
			ProductID: arg.productID,
			Quantity:  arg.quantity,
		})

		if err != nil {
			return err
		}

		//Update the Product quantity

		result.product, err = q.UpdateSoldProduct(ctx, UpdateSoldProductParams{
			Amount: -arg.quantity,
			ID:     arg.productID,
		})

		if err != nil {
			fmt.Printf("%v", err)
			return err
		}

		return nil

	})

	return result, err

}
