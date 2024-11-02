-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS projects(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  user_id TEXT NOT NULL REFERENCES users(id),
  organization_id TEXT NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
  name TEXT NOT NULL,
  key_prefix TEXT NOT NULL,
  type TEXT NOT NULL CHECK(type IN ('global', 'project', 'department')),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS projects_users(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  project_id TEXT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

-- indexes
CREATE UNIQUE INDEX IF NOT EXISTS projects_key ON projects (organization_id, name, key_prefix);
CREATE INDEX IF NOT EXISTS projects_user_key ON projects_users (user_id, project_id);

-- trigger functions
CREATE TRIGGER update_timestamp_trigger_projects
BEFORE UPDATE ON projects
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger_projects_users
BEFORE UPDATE ON projects_users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS projects_users;
DROP TABLE IF EXISTS projects;

DROP TRIGGER IF EXISTS update_timestamp_trigger_projects_users ON projects_users;
DROP TRIGGER IF EXISTS update_timestamp_trigger_projects ON projects;

DROP INDEX IF EXISTS projects_user_key;
DROP INDEX IF EXISTS projects_key;
-- +goose StatementEnd
