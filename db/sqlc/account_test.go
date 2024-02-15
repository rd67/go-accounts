package db

import "testing"


func Test_CreateAccount(t *testing.T) {
	args := CreateAccountParams{
		Name: "Test",
		Balance: 100,
		Currency: "USD",
	}
}