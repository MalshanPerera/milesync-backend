-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS status (
    id CHAR(21) DEFAULT nanoid() PRIMARY KEY,
    name TEXT NOT NULL,
    color TEXT NOT NULL,
    project_id TEXT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    organization_id TEXT NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS labels (
    id CHAR(21) DEFAULT nanoid() PRIMARY KEY,
    name TEXT NOT NULL,
    color TEXT NOT NULL,
    project_id TEXT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    organization_id TEXT NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tasks (
    id CHAR(21) DEFAULT nanoid() PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    user_id TEXT NOT NULL REFERENCES users(id),
    assigner_id TEXT REFERENCES users(id),
    reporter_id TEXT REFERENCES users(id),
    organization_id TEXT NOT NULL REFERENCES organization(id),
    project_id TEXT NOT NULL REFERENCES projects(id),
    status_id TEXT NOT NULL REFERENCES status(id),
    priority SMALLINT NOT NULL DEFAULT 0,
    due_date TIMESTAMP,
    order_index INT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT check_priority CHECK (priority >= 0 AND priority <= 10)
);

CREATE TABLE IF NOT EXISTS task_labels (
    id CHAR(21) DEFAULT nanoid() PRIMARY KEY,
    task_id TEXT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    label_id TEXT NOT NULL REFERENCES labels(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS task_comments (
    id CHAR(21) DEFAULT nanoid() PRIMARY KEY,
    task_id TEXT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    user_id TEXT NOT NULL REFERENCES users(id),
    parent_comment_id TEXT REFERENCES task_comments(id) ON DELETE CASCADE,
    comment TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS task_attachments (
    id CHAR(21) DEFAULT nanoid() PRIMARY KEY,
    task_id TEXT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    user_id TEXT NOT NULL REFERENCES users(id),
    file_name TEXT NOT NULL,
    file_path TEXT NOT NULL,
    file_size BIGINT NOT NULL,
    mime_type TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS task_assignees (
    id CHAR(21) DEFAULT nanoid() PRIMARY KEY,
    task_id TEXT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    user_id TEXT NOT NULL REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS comment_attachments (
    id CHAR(21) DEFAULT nanoid() PRIMARY KEY,
    comment_id TEXT NOT NULL REFERENCES task_comments(id) ON DELETE CASCADE,
    user_id TEXT NOT NULL REFERENCES users(id),
    file_name TEXT NOT NULL,
    file_path TEXT NOT NULL,
    file_size BIGINT NOT NULL,
    mime_type TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE INDEX idx_tasks_user_id ON tasks(user_id);
CREATE INDEX idx_tasks_project_id ON tasks(project_id);
CREATE INDEX idx_tasks_status_id ON tasks(status_id);
CREATE INDEX idx_task_comments_task_id ON task_comments(task_id);
CREATE INDEX idx_task_attachments_task_id ON task_attachments(task_id);
CREATE INDEX idx_task_assignees_task_id ON task_assignees(task_id);

CREATE TRIGGER update_timestamp_trigger_tasks
BEFORE UPDATE ON tasks
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger_task_comments
BEFORE UPDATE ON task_comments
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger_task_attachments
BEFORE UPDATE ON task_attachments
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger_task_assignees
BEFORE UPDATE ON task_assignees
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger_comment_attachments
BEFORE UPDATE ON comment_attachments
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comment_attachments;
DROP TABLE IF EXISTS task_assignees;
DROP TABLE IF EXISTS task_attachments;
DROP TABLE IF EXISTS task_comments;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS status;

DROP TRIGGER IF EXISTS update_timestamp_trigger_comment_attachments ON comment_attachments;
DROP TRIGGER IF EXISTS update_timestamp_trigger_task_assignees ON task_assignees;

DROP TRIGGER IF EXISTS update_timestamp_trigger_task_attachments ON task_attachments;
DROP TRIGGER IF EXISTS update_timestamp_trigger_task_comments ON task_comments;
DROP TRIGGER IF EXISTS update_timestamp_trigger_tasks ON tasks;

ALTER TABLE IF EXISTS tasks DROP CONSTRAINT IF EXISTS check_priority;
-- +goose StatementEnd
