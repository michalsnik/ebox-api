-- +goose Up
CREATE TABLE ebox.users (
    id SERIAL PRIMARY KEY,

    email VARCHAR(50) NOT NULL UNIQUE,
    firstName VARCHAR(50),
    lastName VARCHAR(50),
    avatarUrl VARCHAR,

    createdAt TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE ebox.users
