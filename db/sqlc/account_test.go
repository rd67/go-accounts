package db

import (
	"context"
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
)


func TestCreateAccount(t *testing.T) {

	faker := faker.New()

	args := CreateAccountParams{
		Name: faker.Person().Name(),
		Balance: float64(faker.Currency().Number()),
		Currency: faker.Currency().Currency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.)
}
