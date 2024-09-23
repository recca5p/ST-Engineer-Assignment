package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	UpdateBoardTx(ctx context.Context, args UpdateBoardTxParams) (string, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db)}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	transaction, transactionError := store.db.BeginTx(ctx, nil)

	if transactionError != nil {
		return transactionError
	}

	queries := New(transaction)
	transactionError = fn(queries)

	if transactionError != nil {
		if rollbackErr := transaction.Rollback(); rollbackErr != nil {
			return fmt.Errorf("Transaction error: %v; rollback error: %v", transactionError, rollbackErr)
		}
	}

	return transaction.Commit()
}
