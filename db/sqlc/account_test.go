package db

import (
	"context"
	"testing"
	"github.com/stretchr/testify/require"
	_ "github.com/jaswdr/faker/v2"
)


func TestCreateAccount(t *testing.T) {
	args := CreateAccountParams{
		Name: "Test",
		Balance: 100,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)
}