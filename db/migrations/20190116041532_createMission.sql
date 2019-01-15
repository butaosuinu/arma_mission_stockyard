
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE missions (
    id int NOT NULL AUTO_INCREMENT,
    maker_id int,
    name VARCHAR(255) NOT NULL,
    map_id int,
    mission_type_id int,
    friendly_id int,
    file_url TEXT,
    preset TEXT,
    description TEXT,
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE missions;
