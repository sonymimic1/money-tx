package db

import (
	"context"
	"testing"

	"github.com/sonymimic1/go-transfer/pkg/util"
	"github.com/stretchr/testify/require"
)

func addEntry(t *testing.T, account Account) Entry {

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    int64(util.RandomMoney()),
	}
	got, err := TestQuri.CreateEntry(context.Background(), arg)
	require.NoError(t, err)

	gotID, err := got.LastInsertId()
	require.NoError(t, err)
	entry, err := TestQuri.GetEntry(context.Background(), gotID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestQueries_CreateEntry(t *testing.T) {
	account := addAccount(t)
	addEntry(t, account)
}

func TestQueries_GetEntry(t *testing.T) {
	account := addAccount(t)
	entry := addEntry(t, account)
	entry_result, err := TestQuri.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry_result)

	require.Equal(t, entry_result.AccountID, entry.AccountID)
	require.Equal(t, entry_result.Amount, entry.Amount)
	require.Equal(t, entry_result.ID, entry.ID)
	require.NotZero(t, entry_result.CreatedAt)
}

func TestQueries_ListEntries(t *testing.T) {
	account := addAccount(t)
	for i := 0; i < 10; i++ {
		addEntry(t, account)
	}
	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := TestQuri.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, entry.AccountID, arg.AccountID)
	}
}
