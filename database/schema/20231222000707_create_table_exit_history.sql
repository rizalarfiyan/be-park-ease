-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS exit_history (
    entry_history_id varchar(16) NOT NULL,
    location_code varchar(16) NOT NULL,
    price decimal NOT NULL DEFAULT 0,
    exited_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    exited_by int NOT NULL,
    FOREIGN KEY (location_code) REFERENCES location (code),
    FOREIGN KEY (exited_by) REFERENCES users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE exit_history;
-- +goose StatementEnd
