-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tb_motorcycles(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    year INT NOT NULL,
    model TEXT NOT NULL,
    plate TEXT NOT NULL,
    CONSTRAINT unique_plate UNIQUE (plate)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tb_motorcycles;
-- +goose StatementEnd
