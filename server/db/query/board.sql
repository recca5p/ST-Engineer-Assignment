-- name: CreateBoard :one
INSERT INTO boards (name)
VALUES ($1) RETURNING *;

-- name: GetBoard :one
SELECT *
FROM boards
WHERE id = $1 LIMIT 1;

-- name: ListBoards :many
SELECT *
FROM boards
ORDER BY id LIMIT $1
OFFSET $2;

-- name: UpdateBoard :one
UPDATE boards
SET name = $1, updated_at = NOW()
WHERE id = $2
    RETURNING *;

-- name: DeleteBoard :one
DELETE FROM boards
WHERE id = $1
    RETURNING *;
