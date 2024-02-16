-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateEntry :one
INSERT INTO entries (
  account_id, amount, currency, exchange_rate
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateEntry :one
UPDATE entries
  set account_id = $2, amount = $3, currency = $4, exchange_rate = $5
WHERE id = $1
RETURNING *;

-- name: DeleteEntry :one
UPDATE entries
  set "isDeleted" = true, "updatedAt" = now()
WHERE id = $1
RETURNING *;