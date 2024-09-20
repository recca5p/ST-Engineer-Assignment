-- name: CreateTask :one
INSERT INTO tasks (title, description, column_id, position, due_date)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetTask :one
SELECT *
FROM tasks
WHERE id = $1 LIMIT 1;

-- name: ListTasks :many
SELECT *
FROM tasks
WHERE column_id = $1
ORDER BY position LIMIT $2
OFFSET $3;

-- name: UpdateTask :one
UPDATE tasks
SET title = $1, description = $2, position = $3, due_date = $4, updated_at = NOW()
WHERE id = $5
    RETURNING *;

-- name: DeleteTask :one
DELETE FROM tasks
WHERE id = $1
    RETURNING *;
