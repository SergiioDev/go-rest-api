-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS drivers
(
    id             UUID PRIMARY KEY,
    name           TEXT      NOT NULL,
    lastname         TEXT      NOT NULL,
    team            TEXT      NOT NULL,
    date_birth      TEXT      NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS drivers;