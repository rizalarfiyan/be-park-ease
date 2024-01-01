-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS entry_history (
    id varchars PRIMARY KEY,
    location_code varchar(16) NOT NULL,
    vehicle_type_code varchar(16) NOT NULL,
    vehicle_number varchar(16) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by int NOT NULL,
    FOREIGN KEY (location_code) REFERENCES location (code),
    FOREIGN KEY (vehicle_type_code) REFERENCES vehicle_type (code),
    FOREIGN KEY (created_by) REFERENCES users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE entry_history;
-- +goose StatementEnd
