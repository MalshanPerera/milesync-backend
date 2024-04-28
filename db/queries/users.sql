-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserFromEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateUser :one
UPDATE users set first_name = $2, last_name = $3, email = $4 WHERE id = $1 RETURNING *;