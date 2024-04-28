-- name: CreateOrganization :one
INSERT INTO organization (
  user_id, name, slug
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateOrganization :one
UPDATE organization
SET name = $2, slug = $3
WHERE user_id = $1

RETURNING *;

-- name: GetOrganization :one
SELECT * FROM organization
WHERE user_id = $1 LIMIT 1;

-- name: DeleteOrganization :exec
DELETE FROM organization
WHERE user_id = $1;

-- name: UpdateOrganizationOwner :one
UPDATE organization
SET user_id = $2
WHERE user_id = $1

RETURNING *;