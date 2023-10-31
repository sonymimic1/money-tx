package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {

	defer func() error {
		// transcation 併發情形下 尚未commit前 若有參照的fk被使用會鎖住所以要將此檢查關閉
		// 如同postgreSQL 中的 FOR NO KEY UPDATE.
		if _, err := store.db.ExecContext(ctx, `SET foreign_key_checks = 1`); err != nil {
			return err
		}
		return nil
	}()

	// transcation 併發情形下 尚未commit前 若有參照的fk被使用會鎖住所以要將此檢查關閉
	// 如同postgreSQL 中的 FOR NO KEY UPDATE.
	if _, err := store.db.ExecContext(ctx, `SET foreign_key_checks = 0`); err != nil {
		return err
	}

	// begin Tx.
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// doing
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rberr := tx.Rollback(); rberr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rberr)
		}
		return err
	}

	// commit Tx.
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:amount`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

var txKey = struct{}{}

func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {

		// 處理交易明細
		createTransferResult, err := q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: uint64(arg.FromAccountID),
			ToAccountID:   uint64(arg.ToAccountID),
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		transferID, err := createTransferResult.LastInsertId()
		if err != nil {
			return err
		}

		transfer, err := q.GetTransfer(ctx, transferID)
		if err != nil {
			return err
		}
		result.Transfer = transfer

		// 處理匯出人明目
		fromEntryResult, err := q.CreateEntry(ctx, CreateEntryParams{
			AccountID: uint64(arg.FromAccountID),
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		fromEntryID, err := fromEntryResult.LastInsertId()
		if err != nil {
			return err
		}

		fromEntry, err := q.GetEntry(ctx, fromEntryID)
		if err != nil {
			return err
		}
		result.FromEntry = fromEntry

		// 處理匯入人明目
		toEntryResult, err := q.CreateEntry(ctx, CreateEntryParams{
			AccountID: uint64(arg.ToAccountID),
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}
		toEntryID, err := toEntryResult.LastInsertId()
		if err != nil {
			return err
		}

		toEntry, err := q.GetEntry(ctx, toEntryID)
		if err != nil {
			return err
		}
		result.ToEntry = toEntry

		//TODO: update account's balance
		if arg.FromAccountID < arg.ToAccountID {

			err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     uint64(arg.FromAccountID),
				Amount: -arg.Amount,
			})
			if err != nil {
				return err
			}

			account1, err := q.GetAccount(ctx, uint64(arg.FromAccountID))
			if err != nil {
				return err
			}
			result.FromAccount = account1

			err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     uint64(arg.ToAccountID),
				Amount: arg.Amount,
			})
			if err != nil {
				return err
			}

			account2, err := q.GetAccount(ctx, uint64(arg.ToAccountID))
			if err != nil {
				return err
			}
			result.ToAccount = account2
		} else {

			err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     uint64(arg.ToAccountID),
				Amount: arg.Amount,
			})
			if err != nil {
				return err
			}

			account2, err := q.GetAccount(ctx, uint64(arg.ToAccountID))
			if err != nil {
				return err
			}
			result.ToAccount = account2

			err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     uint64(arg.FromAccountID),
				Amount: -arg.Amount,
			})
			if err != nil {
				return err
			}

			account1, err := q.GetAccount(ctx, uint64(arg.FromAccountID))
			if err != nil {
				return err
			}
			result.FromAccount = account1
		}

		return nil
	})

	return result, err
}
