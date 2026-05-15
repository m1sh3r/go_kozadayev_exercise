-- name: CreateUser :one
INSERT INTO users (name) VALUES ($1) RETURNING *;

-- name: AddDocument :one
INSERT INTO documents (title, user_id) VALUES ($1, $2) RETURNING *;

-- name: GetUserDocuments :many
SELECT * FROM documents WHERE user_id = $1;
