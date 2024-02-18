-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CountAccounts :one
SELECT COUNT(*) FROM accounts;

-- name: CreateAccount :one
INSERT INTO accounts (
  "name", "balance", "currency"
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateAccount :one
UPDATE accounts
  set "balance" = $2, "updatedAt" = now()
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :one
UPDATE accounts
  set "isDeleted" = true, "updatedAt" = now()
WHERE id = $1
RETURNING *;