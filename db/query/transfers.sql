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
  sender_account_id, receiver_account_id, amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateTransfer :one
UPDATE transfers
  set sender_account_id = $2, receiver_account_id= $3, amount = $4
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :one
UPDATE transfers
  set "isDeleted" = true, "updatedAt" = now()
WHERE id = $1
RETURNING *;