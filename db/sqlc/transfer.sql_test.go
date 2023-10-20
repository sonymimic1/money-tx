package db

import (
	"context"
	"testing"

	"github.com/sonymimic1/go-transfer/pkg/util"
	"github.com/stretchr/testify/require"
)

func addTransfer(t *testing.T, account1 Account, account2 Account) Transfer {

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	transfer, err := TestQuri.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)

	transferID, err := transfer.LastInsertId()
	require.NoError(t, err)
	transfer_result, err := TestQuri.GetTransfer(context.Background(), transferID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer_result)

	require.Equal(t, transfer_result.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer_result.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfer_result.Amount, arg.Amount)

	require.NotZero(t, transfer_result.ID)
	require.NotZero(t, transfer_result.CreatedAt)
	return transfer_result
}
func TestQueries_CreateTransfer(t *testing.T) {
	account1 := addAccount(t)
	account2 := addAccount(t)
	addTransfer(t, account1, account2)
}

func TestQueries_GetTransfer(t *testing.T) {
	account1 := addAccount(t)
	account2 := addAccount(t)
	addTransfer(t, account1, account2)
}

func TestQueries_ListTransfers(t *testing.T) {
	account1 := addAccount(t)
	account2 := addAccount(t)

	for i := 0; i < 10; i++ {
		addTransfer(t, account1, account2)
		addTransfer(t, account2, account1)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}
	transfers, err := TestQuri.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}
}
