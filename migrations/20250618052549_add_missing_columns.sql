-- +goose Up
-- +goose StatementBegin
ALTER TABLE todos
    ADD COLUMN IF NOT EXISTS due_date TIMESTAMP,
    ADD COLUMN IF NOT EXISTS groupName TEXT,
    ADD COLUMN IF NOT EXISTS description TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE todos
    DROP COLUMN IF EXISTS due_date,
    DROP COLUMN IF EXISTS groupName,
    DROP COLUMN IF EXISTS description;
-- +goose StatementEnd
