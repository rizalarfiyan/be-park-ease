-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS setting (
    key varchar(100) NOT NULL,
    value varchar(500) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE setting;
-- +goose StatementEnd
