-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS organization_users(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  organization_id TEXT NOT NULL REFERENCES organization(id) ON DELETE CASCADE,
  user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS organization_users_key ON organization_users(user_id,organization_id);

CREATE TRIGGER update_timestamp_trigger_organization_users
BEFORE UPDATE ON organization_users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS organization_users;
DROP TRIGGER IF EXISTS update_timestamp_trigger_organization_users ON organization_users;
-- +goose StatementEnd
