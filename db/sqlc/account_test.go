package db

import (
	"context"
	"testing"
	"github.com/stretchr/testify/require"
)


func Test_CreateAccount(t *testing.T) {
	args := CreateAccountParams{
		Name: "Test",
		Balance: 100,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)
}