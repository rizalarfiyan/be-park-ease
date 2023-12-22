-- +goose Up
CREATE TYPE user_role AS ENUM ('admin', 'karyawan');

CREATE TYPE user_status AS ENUM ('active', 'banned');

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL,
    username varchar(20) UNIQUE NOT NULL,
    password varchar(255) NOT NULL,
    role user_role NOT NULL DEFAULT ('karyawan'),
    status user_status NOT NULL DEFAULT ('active'),
    token varchar(255),
    expired_at timestamp,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp
);

-- +goose Down
DROP TABLE users;
DROP TYPE user_role;
DROP TYPE user_status;
