-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY id LIMIT ?, ?;

-- name: CreateAccount :execresult
INSERT INTO accounts (name, balance, currency) VALUES ($1, $2, $3);

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;

-- name: UpdateAccountDetails :execresult
UPDATE accounts SET name = $2 WHERE id = $1;

-- name: UpdateAccountBalance :execresult
UPDATE accounts SET balance = balance + $2 WHERE id = $1;