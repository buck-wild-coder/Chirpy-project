-- +goose Up
ALTER TABLE users
    ADD COLUMN hashed_password text default 'unset' not NULL;


-- +goose down
ALTER TABLE users
    DROP COLUMN IF EXISTS hashed_password;