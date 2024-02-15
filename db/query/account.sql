-- name: GetAccount :one
SELECT * FROM accounts WHERE id = ? LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY id LIMIT ?, ?;

-- name: CreateAccount :execresult
INSERT INTO accounts (name, balance, currency) VALUES (?, ?, ?);

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = ?;

-- name: UpdateAccountDetails :execresult
UPDATE accounts SET name = ? WHERE id = ?;

-- name: UpdateAccountBalance :execresult
UPDATE accounts SET balance = balance + ? WHERE id = ?;