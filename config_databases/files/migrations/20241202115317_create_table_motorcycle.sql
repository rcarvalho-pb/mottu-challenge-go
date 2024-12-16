-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tb_motorcycles(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    year INT NOT NULL,
    model TEXT NOT NULL,
    plate TEXT NOT NULL,
    created_at TIMESTAMP DEFAUL CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_located BOOLEAN DEFAULT FALSE,
    CONSTRAINT unique_plate UNIQUE (plate),
    CONSTRAINT fk_userid FOREIGN KEY (user_id) REFERENCES tb_users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tb_motorcycles;
-- +goose StatementEnd
