-- name: CreateTodo :one
INSERT INTO todos (
  title, 
  note,
  due_date
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTodo :one
UPDATE todos
  set title = $2,
  note = $3, 
  due_date = $4  
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;

