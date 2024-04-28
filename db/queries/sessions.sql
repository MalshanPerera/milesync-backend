-- name: CreateSession :one
INSERT INTO sessions (
  user_id, "access_token", "refresh_token", "expires_at"
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateSession :one
UPDATE sessions
SET "access_token" = $2, "refresh_token" = $3, "expires_at" = $4
WHERE user_id = $1
RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions
WHERE user_id = $1 LIMIT 1;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE user_id = $1
RETURNING *;
