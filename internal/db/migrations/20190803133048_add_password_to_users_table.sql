-- +goose Up
ALTER TABLE ebox.users ADD COLUMN password VARCHAR(60);

-- +goose Down
ALTER TABLE ebox.users DROP COLUMN password;
