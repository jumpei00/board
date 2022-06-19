-- +migrate Up

CREATE TABLE IF NOT EXISTS comments (
    `comment_key` VARCHAR(128) NOT NULL,
    `thread_key` VARCHAR(128) NOT NULL,
    `contributor` VARCHAR(128) NOT NULL,
    `comment` VARCHAR(2048) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`comment_key`),
    FOREIGN KEY `fk_thread_key` (`thread_key`) REFERENCES `threads` (`thread_key`)
);

-- +migrate Down

DROP TABLE IF EXISTS comments;