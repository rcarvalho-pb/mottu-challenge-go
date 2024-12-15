-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tb_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    name TEXT NOT NULL,
    birth_date TIMESTAMP NOT NULL,
    cnpj TEXT NOT NULL,
    cnh TEXT NOT NULL,
    cnh_type TEXT NOT NULL,
    cnh_file_path TEXT,
    active_location BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_user_cnpj UNIQUE (cnpj),
    CONSTRAINT unique_user_cnh UNIQUE (cnh),
    active BOOLEAN DEFAULT TRUE
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tb_users;
-- +goose StatementEnd
