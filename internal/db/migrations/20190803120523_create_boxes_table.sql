-- +goose Up
CREATE TABLE ebox.boxes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE ebox.boxes
