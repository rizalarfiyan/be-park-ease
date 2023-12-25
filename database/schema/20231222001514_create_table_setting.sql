-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS setting (
    key varchar(100) NOT NULL UNIQUE PRIMARY KEY,
    value varchar(500) NOT NULL,
    description VARCHAR(255) NOT NULL DEFAULT ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE setting;
-- +goose StatementEnd
