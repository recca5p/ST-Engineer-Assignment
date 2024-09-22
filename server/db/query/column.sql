-- name: CreateColumn :one
INSERT INTO columns (name, board_id, position)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetColumn :one
SELECT *
FROM columns
WHERE id = $1 LIMIT 1;

-- name: GetColumnByPosition :one
SELECT *
FROM columns
WHERE position = $1 LIMIT 1;

-- name: ListColumns :many
SELECT *
FROM columns
WHERE board_id = $1
ORDER BY position LIMIT $2
OFFSET $3;

-- name: UpdateColumn :one
UPDATE columns
SET name = $1, position = $2, updated_at = NOW()
WHERE id = $3
    RETURNING *;

-- name: DeleteColumn :one
DELETE FROM columns
WHERE id = $1
    RETURNING *;
