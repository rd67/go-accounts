package db

import (
	"context"
	"testing"

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