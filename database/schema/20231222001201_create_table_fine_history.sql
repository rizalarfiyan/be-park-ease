-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS fine_history (
    entry_history_id varchar(16) NOT NULL,
    location_code varchar(16) NOT NULL,
    price decimal NOT NULL DEFAULT 0,
    identity varchar(32) NOT NULL,
    vehicle_identity varchar(32) NOT NULL,
    name varchar(100) NOT NULL,
    address varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    fined_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fined_by int NOT NULL,
    FOREIGN KEY (location_code) REFERENCES location (code),
    FOREIGN KEY (fined_by) REFERENCES users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE fine_history;
-- +goose StatementEnd
