-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateAccount :one
INSERT INTO accounts (
  name, balance, currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateAccountDetails :one
UPDATE accounts
  set name = $2
WHERE id = $1
RETURNING *;

-- name: UpdateAccountBalance :one
UPDATE accounts
  set balance = balance + $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;