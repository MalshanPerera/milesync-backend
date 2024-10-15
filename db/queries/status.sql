-- name: CreateStatus :one
INSERT INTO status (
  name, color, project_id, organization_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateStatus :one
UPDATE status
SET name = $2, color = $3, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteStatus :exec
DELETE FROM status
WHERE id = $1 AND project_id = $2 AND organization_id = $3;

-- name: GetStatus :one
SELECT * FROM status
WHERE id = $1;

-- name: GetStatuses :many
SELECT * FROM status
WHERE project_id = $1 AND organization_id = $2;