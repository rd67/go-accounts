// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (sql.Result, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (sql.Result, error)
	CreateTransfer(ctx context.Context) (sql.Result, error)
	DeleteAccount(ctx context.Context, id uint64) error
	DeleteEntry(ctx context.Context, id uint64) error
	DeleteTransfer(ctx context.Context) error
	GetAccount(ctx context.Context, id uint64) (Account, error)
	GetEntry(ctx context.Context, id uint64) (Entry, error)
	GetTransfer(ctx context.Context) (Transfer, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	UpdateAccountBalance(ctx context.Context, arg UpdateAccountBalanceParams) (sql.Result, error)
	UpdateAccountDetails(ctx context.Context, arg UpdateAccountDetailsParams) (sql.Result, error)
	UpdateEntry(ctx context.Context, arg UpdateEntryParams) (sql.Result, error)
	UpdateTransfer(ctx context.Context) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)
