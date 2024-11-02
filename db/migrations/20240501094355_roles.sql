-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  organization_id TEXT NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
  default_id TEXT,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  permissions TEXT[] NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS roles_projects(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  role_id TEXT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
  project_id TEXT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

-- indexes
CREATE UNIQUE INDEX IF NOT EXISTS roles_organization_key ON roles (organization_id, name);

CREATE TRIGGER update_timestamp_trigger_roles
BEFORE UPDATE ON roles
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger_roles_projects
BEFORE UPDATE ON roles_projects
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles_projects;
DROP TABLE IF EXISTS roles;

DROP TRIGGER IF EXISTS update_timestamp_trigger_roles_projects ON roles_projects;
DROP TRIGGER IF EXISTS update_timestamp_trigger_roles ON roles;
-- +goose StatementEnd
