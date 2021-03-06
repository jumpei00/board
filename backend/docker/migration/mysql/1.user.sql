-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    `id` VARCHAR(128) NOT NULL,
    `username` VARCHAR(128) NOT NULL,
    `password` VARBINARY(128) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);

-- +migrate Down

DROP TABLE IF EXISTS users;