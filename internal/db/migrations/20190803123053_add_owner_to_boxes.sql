-- +goose Up
ALTER TABLE ebox.boxes ADD COLUMN ownerID INTEGER REFERENCES ebox.users;

-- +goose Down
ALTER TABLE ebox.boxes DROP COLUMN ownerID;
