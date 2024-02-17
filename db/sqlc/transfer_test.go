package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	args := CreateTransferParams{
		SenderAccountID:   account1.ID,
		ReceiverAccountID: account2.ID,
		Amount:            int64(faker.RandomNumber(5)),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, args.SenderAccountID, transfer.SenderAccountID)
	require.Equal(t, args.ReceiverAccountID, transfer.ReceiverAccountID)
	require.Equal(t, args.Amount, transfer.Amount)
	require.NotNil(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	createRandomTransfer(t, account1, account2)

}
