-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserFromEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateUser :one
UPDATE users set name = $2, email = $3 WHERE id = $1 RETURNING *;