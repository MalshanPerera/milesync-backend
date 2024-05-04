-- name: CreateProject :one
INSERT INTO projects (
  user_id, organization_id, name, key_prefix, type
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateProject :one
UPDATE projects
SET
  name = $2,
  key_prefix = $3
WHERE
  id = $1 AND user_id = $4
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE id = $1 AND user_id = $2;

-- name: GetProject :one
SELECT * FROM projects
WHERE id = $1 AND user_id = $2 LIMIT 1;

-- name: GetProjectKeyPrefixUsed :one
SELECT EXISTS (
    SELECT 1
    FROM projects
    WHERE projects.key_prefix = $1
);

-- name: GetProjects :many
SELECT * FROM projects
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetProjectByKeyPrefix :one
SELECT * FROM projects
WHERE key_prefix = $1
LIMIT 1;
