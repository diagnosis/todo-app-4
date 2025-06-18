-- +goose Up
-- +goose StatementBegin
CREATE TABLE todos (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       completed BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       due_date TIMESTAMP,
                       groupName TEXT,
                       description TEXT
);

CREATE OR REPLACE FUNCTION update_updated_at()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_timestamp
    BEFORE UPDATE ON todos
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_timestamp ON todos;
DROP FUNCTION IF EXISTS update_updated_at;
DROP TABLE todos;
-- +goose StatementEnd