-- name: GetEntry :one
SELECT * FROM entries WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries ORDER BY id LIMIT ?, ?;

-- name: CreateEntry :execresult
INSERT INTO entries (account_id, amount, currency, exchange_rate) VALUES ($1, $2, $3, $4);

-- name: DeleteEntry :exec
DELETE FROM entries WHERE id = $1;

-- name: UpdateEntry :execresult
UPDATE entries SET account_id = $2, amount = $3, currency = $4, exchange_rate = $5 WHERE id = $1;