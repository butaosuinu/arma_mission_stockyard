
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE maps (
    id int NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    steam_url TEXT,
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE maps;
