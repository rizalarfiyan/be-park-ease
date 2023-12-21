-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS location (
    code varchar(16) PRIMARY KEY,
    name varchar(255) NOT NULL,
    is_exit boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by int NOT NULL,
    updated_at timestamp,
    updated_by int,
    deleted_at timestamp,
    deleted_by int,
    FOREIGN KEY (created_by) REFERENCES users (id),
    FOREIGN KEY (updated_by) REFERENCES users (id),
    FOREIGN KEY (deleted_by) REFERENCES users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE location;
-- +goose StatementEnd
