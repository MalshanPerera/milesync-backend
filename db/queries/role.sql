-- name: CreateRole :one
INSERT INTO roles (
  organization_id, default_id, name, description, permissions
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;


-- name: GetRole :one
SELECT * FROM roles
WHERE organization_id = $1 LIMIT 1;

-- name: GetRolesForOrganization :many
SELECT * FROM roles
WHERE organization_id = $1 ORDER BY created_at DESC;

-- name: GetRolesForProject :many
SELECT * FROM roles_projects
WHERE project_id = $1;

-- name: GetRolesForUser :many
SELECT roles.* FROM roles_projects
JOIN roles on roles.id = roles_projects.role_id
WHERE roles_projects.user_id = $1 AND roles.organization_id = $2;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE organization_id = $1;

-- name: UpdateRole :one
UPDATE roles
SET default_id = $2, name = $3, description = $4, permissions = $5
WHERE id = $1 AND organization_id = $6
RETURNING *;

-- name: AddPermissionToRole :one
UPDATE roles
SET permissions = array_append(permissions, $3)
WHERE organization_id = $1 AND id = $2
RETURNING *;

-- name: RemovePermissionFromRole :exec
UPDATE roles
SET permissions = array_remove(permissions, $3)
WHERE organization_id = $1 AND id = $2;


-- name: AddRoleToUser :one
INSERT INTO roles_projects (
  project_id, role_id, user_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: RemoveRoleFromUser :exec
DELETE FROM roles_projects
WHERE project_id = $1 AND role_id = $2 AND user_id = $3;