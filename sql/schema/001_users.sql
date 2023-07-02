-- +goose Up
CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    name text NOT NULL
);

-- +goose Down
DROP TABLE users;
