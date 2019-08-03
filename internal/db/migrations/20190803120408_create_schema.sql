-- +goose Up
CREATE SCHEMA ebox;

-- +goose Down
DROP SCHEMA ebox RESTRICT;
