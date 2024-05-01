-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS projects(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  organization_id TEXT NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
  name UNIQUE TEXT NOT NULL,
  key_prefix UNIQUE TEXT NOT NULL,
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
CREATE UNIQUE INDEX ON projects_key (organization_id, name, key_prefix);
CREATE UNIQUE INDEX ON projects_users_key (organization_id, project_id, user_id);

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
DROP TRIGGER IF EXISTS update_timestamp_trigger_projects ON "projects";
DROP TABLE IF EXISTS "projects";

DROP TRIGGER IF EXISTS update_timestamp_trigger_projects_users ON "projects_users";
DROP TABLE IF EXISTS "projects_users";

DROP INDEX IF EXISTS projects_key;
DROP INDEX IF EXISTS projects_users_key;
-- +goose StatementEnd
