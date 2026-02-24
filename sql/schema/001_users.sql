-- +goose Up
CREATE TABLE users (
    id int primary key,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    email text unique not null
);

-- +goose Down
DROP TABLE users;