-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  permissions TEXT[] NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS roles_users(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  role_id TEXT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
  user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS roles_projects(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  role_id TEXT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
  project_id TEXT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS roles_organization(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  role_id TEXT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
  organization_id TEXT NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

CREATE UNIQUE INDEX ON roles_users (role_id, user_id);
CREATE UNIQUE INDEX ON roles_projects (role_id, project_id);
CREATE UNIQUE INDEX ON roles_organization (role_id, organization_id);

CREATE TRIGGER update_timestamp_trigger_roles
BEFORE UPDATE ON roles
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger_roles_users
BEFORE UPDATE ON "roles_users"
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger_roles_projects
BEFORE UPDATE ON "roles_projects"
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_timestamp_trigger_roles_organization
BEFORE UPDATE ON "roles_organization"
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_timestamp_trigger_roles ON "roles";
DROP TABLE IF EXISTS roles;

DROP TRIGGER IF EXISTS update_timestamp_trigger_roles_users ON "roles_users";
DROP TABLE IF EXISTS "roles_users";

DROP TRIGGER IF EXISTS update_timestamp_trigger_roles_projects ON "roles_projects";
DROP TABLE IF EXISTS "roles_projects";

DROP TRIGGER IF EXISTS update_timestamp_trigger_roles_organization ON "roles_organization";
DROP TABLE IF EXISTS "roles_organization";
-- +goose StatementEnd
