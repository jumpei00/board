-- +migrate Up

CREATE TABLE IF NOT EXISTS visitors (
    `id` INTEGER NOT NULL AUTO_INCREMENT,
    `yesterday_visitor` INTEGER NOT NULL,
    `today_visitor` INTEGER NOT NULL,
    `visitor_sum` INTEGER NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);
INSERT INTO visitors VALUES (1, 0, 0, 0, NOW(), NOW());

-- +migrate Down

DROP TABLE IF EXISTS visitors;