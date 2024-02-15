-- name: GetEntry :one
SELECT * FROM entries WHERE id = ? LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries ORDER BY id LIMIT ?, ?;

-- name: CreateEntry :execresult
INSERT INTO entries (account_id, amount, currency, exchange_rate) VALUES (?, ?, ?, ?);

-- name: DeleteEntry :exec
DELETE FROM entries WHERE id = ?;

-- name: UpdateEntry :execresult
UPDATE entries SET account_id = ?, amount = ?, currency = ?, exchange_rate = ? WHERE id = ?;