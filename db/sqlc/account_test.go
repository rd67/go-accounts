package db

import (
	"context"
	"testing"
	"time"

	fk "github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
)

var faker = fk.New()

func createTestAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Name: faker.Person().Name(),
		Balance: int64(faker.RandomNumber(5)),
		Currency: faker.Currency().Code(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Name, args.Name)
	require.Equal(t, account.Balance, args.Balance)
	require.Equal(t, account.Currency, args.Currency)

	require.NotEqual(t, 0, account.ID)
	require.NotNil(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createTestAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createTestAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createTestAccount(t)

	account2, err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.IsDeleted, true)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createTestAccount(t)


	args := UpdateAccountParams{
		ID: account1.ID,
		Balance: int64(faker.RandomNumber(5)),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.UpdatedAt, account2.UpdatedAt, time.Second)

	require.Equal(t, account2.Balance, args.Balance)
}