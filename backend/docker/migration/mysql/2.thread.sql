-- +migrate Up

CREATE TABLE IF NOT EXISTS threads (
    `thread_key` VARCHAR(128) NOT NULL,
    `title` VARCHAR(256) NOT NULL,
    `contributor` VARCHAR(128) NOT NULL,
    `views` INTEGER NOT NULL,
    `comment_sum` INTEGER NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`thread_key`)
);

-- +migrate Down

DROP TABLE IF EXISTS threads;