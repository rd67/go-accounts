-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateTransfer :one
INSERT INTO transfers (
  sender_account_id, receiver_account_id, amount, currency, exchange_rate, status
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateTransfer :one
UPDATE transfers
  set sender_account_id = $2, receiver_account_id= $3, amount = $4, currency = $5, exchange_rate = $6, status = $7
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;