package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/sonymimic1/go-transfer/pkg/util"
	"github.com/stretchr/testify/require"
)

func addAccount(t *testing.T) Account {

	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	var account Account

	got, err := TestQuri.CreateAccount(context.Background(), arg)
	require.NoError(t, err)

	afterInsertID, _ := got.LastInsertId()
	account, _ = TestQuri.GetAccount(context.Background(), uint64(afterInsertID))

	require.NoError(t, err)
	require.NotEmpty(t, got)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestQueries_CreateAccount(t *testing.T) {
	addAccount(t)
}

func TestQueries_UpdateAccount(t *testing.T) {

	account1 := addAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}
	err := TestQuri.UpdateAccount(context.Background(), arg)
	if err != nil {
		fmt.Printf("UpdateAccount Erro \n")
	}
	account1_update, _ := TestQuri.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account1_update)

	require.Equal(t, account1.ID, account1_update.ID)
	require.Equal(t, account1.Owner, account1_update.Owner)
	require.Equal(t, arg.Balance, account1_update.Balance)
	require.Equal(t, account1.Currency, account1_update.Currency)
	require.Equal(t, account1.CreatedAt, account1_update.CreatedAt)

}

func TestQueries_DeleteAccount(t *testing.T) {
	account1 := addAccount(t)
	err := TestQuri.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account1_delete, err := TestQuri.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account1_delete)
}

func TestQueries_ListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		addAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := TestQuri.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
