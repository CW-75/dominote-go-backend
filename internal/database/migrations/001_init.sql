-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS game_tables (
    uid UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    best_score INT DEFAULT 0,
    player_with_most_victories VARCHAR(255)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF NOT EXISTS game_tables;
