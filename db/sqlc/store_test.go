package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStore_TransferTx(t *testing.T) {
	store := NewStore(TConn)

	account1 := addAccount(t)
	account2 := addAccount(t)
	fmt.Println(">>before:", account1.Balance, account2.Balance)
	n := 5

	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	//併發處理多筆資料
	for i := 0; i < n; i++ {
		txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: int64(account1.ID),
				ToAccountID:   int64(account2.ID),
				Amount:        amount,
			})
			if err != nil {
				fmt.Printf("err:%v", err)
			}
			errs <- err
			results <- result
		}()
	}

	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		//check any error
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		//check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, account1.ID)
		require.Equal(t, transfer.ToAccountID, account2.ID)
		require.Equal(t, transfer.Amount, amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		//check transfer
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, fromEntry.AccountID, account1.ID)
		require.Equal(t, fromEntry.Amount, -amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, toEntry.AccountID, account2.ID)
		require.Equal(t, toEntry.Amount, amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		//check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, fromAccount.ID, account1.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, toAccount.ID, account2.ID)

		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := diff1 / amount
		require.True(t, k >= 1 && k <= int64(n))
		existed[int(k)] = true
	}

	updateAccount1, err := TestQuri.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updateAccount2, err := TestQuri.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">>before:", updateAccount1.Balance, updateAccount2.Balance)

	require.Equal(t, account1.Balance-int64(n)*amount, updateAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updateAccount2.Balance)
}

func TestStore_TransferTxDeadLock(t *testing.T) {
	store := NewStore(TConn)

	account1 := addAccount(t)
	account2 := addAccount(t)
	fmt.Println(">>before:", account1.Balance, account2.Balance)
	n := 10

	amount := int64(10)

	errs := make(chan error)

	//併發處理多筆資料
	for i := 0; i < n; i++ {

		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}
		txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)
			_, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: int64(fromAccountID),
				ToAccountID:   int64(toAccountID),
				Amount:        amount,
			})
			if err != nil {
				fmt.Printf("err:%v", err)
			}
			errs <- err

		}()
	}

	for i := 0; i < n; i++ {
		//check any error
		err := <-errs
		require.NoError(t, err)
	}

	updateAccount1, err := TestQuri.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updateAccount2, err := TestQuri.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">>before:", updateAccount1.Balance, updateAccount2.Balance)

	require.Equal(t, account1.Balance, updateAccount1.Balance)
	require.Equal(t, account2.Balance, updateAccount2.Balance)
}
