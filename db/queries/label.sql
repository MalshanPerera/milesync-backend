-- name: CreateLabel :one
INSERT INTO labels (
  name, color, project_id, organization_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateLabel :one
UPDATE labels
SET name = $2, color = $3, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteLabel :exec
DELETE FROM labels
WHERE id = $1 AND project_id = $2 AND organization_id = $3;

-- name: GetLabel :one
SELECT * FROM labels
WHERE id = $1 LIMIT 1;

-- name: GetLabels :many
SELECT * FROM labels
WHERE project_id = $1 AND organization_id = $2;