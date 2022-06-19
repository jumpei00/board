-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    `id` INTEGER NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(128) NOT NULL,
    `password` VARBINARY(128) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);

-- +migrate Down

DROP TABLE IF EXISTS users;