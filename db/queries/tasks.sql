-- name: CreateTask :one
INSERT INTO tasks (
  user_id, assigner_id, reporter_id, organization_id, project_id, status_id, title, description, priority, due_date, order_index
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks
SET user_id = $2, assigner_id = $3, reporter_id = $4, organization_id = $5, project_id = $6, status_id = $7,
    title = $8, description = $9, priority = $10, due_date = $11, order_index = $12, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTask :exec
UPDATE tasks
SET deleted_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: GetTasks :many
SELECT * FROM tasks
WHERE organization_id = $1 AND project_id = $2 AND deleted_at IS NULL
ORDER BY order_index ASC, created_at DESC;

-- name: CreateTaskComment :one
INSERT INTO task_comments (
  task_id, user_id, parent_comment_id, comment
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateTaskComment :one
UPDATE task_comments
SET comment = $2, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTaskComment :exec
DELETE FROM task_comments
WHERE id = $1;

-- name: GetTaskComments :many
SELECT * FROM task_comments
WHERE task_id = $1
ORDER BY created_at ASC;

-- name: CreateTaskAttachment :one
INSERT INTO task_attachments (
  task_id, user_id, file_name, file_path, file_size, mime_type
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: DeleteTaskAttachment :exec
DELETE FROM task_attachments
WHERE id = $1;

-- name: GetTaskAttachments :many
SELECT * FROM task_attachments
WHERE task_id = $1
ORDER BY created_at DESC;

-- name: CreateTaskAssignee :one
INSERT INTO task_assignees (
  task_id, user_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteTaskAssignee :exec
DELETE FROM task_assignees
WHERE task_id = $1 AND user_id = $2;

-- name: GetTaskAssignees :many
SELECT * FROM task_assignees
WHERE task_id = $1
ORDER BY created_at ASC;

-- name: CreateCommentAttachment :one
INSERT INTO comment_attachments (
  comment_id, user_id, file_name, file_path, file_size, mime_type
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: DeleteCommentAttachment :exec
DELETE FROM comment_attachments
WHERE id = $1;

-- name: GetCommentAttachments :many
SELECT * FROM comment_attachments
WHERE comment_id = $1
ORDER BY created_at ASC;