package db

import (
	"context"
	"testing"
	"time"

	"strings"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {

	account := createRandomAccount(t)

	args := CreateEntryParams{
		AccountID:    account.ID,
		Amount:       int64(faker.RandomNumber(5)),
		Currency:     strings.Trim(faker.Currency().Code(), " "),
		ExchangeRate: 1,
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)
	require.Equal(t, args.Currency, entry.Currency)
	require.Equal(t, args.ExchangeRate, entry.ExchangeRate)

	require.NotEqual(t, entry.ID, 0)
	require.NotNil(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestDeleteEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	entry2, err := testQueries.DeleteEntry(context.Background(), entry1.ID)

	require.NoError(t, err)

	require.Equal(t, entry2.IsDeleted, true)
	require.NotEqual(t, entry1.UpdatedAt, entry2.UpdatedAt)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.Currency, entry2.Currency)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	args := ListEntriesParams{
		Limit: 5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, entries, int(args.Limit))

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.NotZero(t, entry.ID)
	}
}

func TestUpdateEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	args := UpdateEntryParams{
			ID: entry1.ID,
			AccountID: entry1.AccountID,
			Amount: int64(faker.RandomNumber(5)),
			Currency: entry1.Currency,
			ExchangeRate: 2,
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, args.AccountID, entry2.AccountID)
	require.Equal(t, args.Amount, entry2.Amount)
	require.Equal(t, args.Currency, entry2.Currency)
	require.Equal(t, args.ExchangeRate, entry2.ExchangeRate)

	require.WithinDuration(t, entry1.UpdatedAt, entry2.UpdatedAt, time.Second)
}