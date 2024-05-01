-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS organization(
  id char(21) DEFAULT nanoid() PRIMARY KEY,
  user_id TEXT NOT NULL REFERENCES users(id),
  name TEXT NOT NULL,
  slug TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP
);
CREATE TRIGGER update_timestamp_trigger BEFORE UPDATE ON "organization" FOR EACH ROW EXECUTE FUNCTION update_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_timestamp_trigger ON "organization";
DROP TABLE IF EXISTS "organization";
-- +goose StatementEnd
