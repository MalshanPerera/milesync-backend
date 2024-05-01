-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS projects(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  user_id TEXT NOT NULL REFERENCES users(id),
  organization_id TEXT NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
  name TEXT NOT NULL UNIQUE,
  key_prefix TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS projects_organization(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  organization_id TEXT NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
  project_id TEXT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS projects_users(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  organization_id TEXT NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
  project_id TEXT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

-- indexes
CREATE UNIQUE INDEX IF NOT EXISTS projects_key ON projects (organization_id, name, key_prefix);
CREATE INDEX IF NOT EXISTS projects_user_key ON projects_users (user_id, organization_id, project_id);
CREATE UNIQUE INDEX IF NOT EXISTS projects_organization_key ON projects_organization (organization_id, project_id);

-- trigger functions
CREATE TRIGGER update_timestamp_trigger_projects
BEFORE UPDATE ON projects
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger_projects_users
BEFORE UPDATE ON projects_users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger_projects_organization
BEFORE UPDATE ON projects_organization
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS projects_users;
DROP TABLE IF EXISTS projects_organization;
DROP TABLE IF EXISTS projects;

DROP TRIGGER IF EXISTS update_timestamp_trigger_projects_users ON projects_users;
DROP TRIGGER IF EXISTS update_timestamp_trigger_projects_organization ON projects_organization;
DROP TRIGGER IF EXISTS update_timestamp_trigger_projects ON projects;

DROP INDEX IF EXISTS projects_user_key;
DROP INDEX IF EXISTS projects_organization_key;
DROP INDEX IF EXISTS projects_key;
-- +goose StatementEnd
