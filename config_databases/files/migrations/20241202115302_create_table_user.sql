-- +goose Up
-- +goose StatementBegin
/*CREATE TABLE IF NOT EXISTS tb_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    role INTEGER DEFAULT 1,
    name TEXT NOT NULL,
    birth_date TIMESTAMP NOT NULL,
    cnpj TEXT NOT NULL,
    cnh TEXT NOT NULL,
    cnh_type TEXT NOT NULL,
    cnh_file_path TEXT NOT NULL DEFAULT '',
    active_location BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    active BOOLEAN DEFAULT TRUE,
    CONSTRAINT unique_user_cnpj UNIQUE (cnpj),
    CONSTRAINT unique_user_cnh UNIQUE (cnh)
    );
*/
    CREATE TABLE IF NOT EXISTS tb_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    role INTEGER DEFAULT 1,
    name TEXT NOT NULL,
    birth_date TIMESTAMP NOT NULL,  -- Alterado para DATE
    cnpj TEXT NOT NULL,
    cnh TEXT NOT NULL,
    cnh_type TEXT NOT NULL,
    cnh_file_path TEXT NOT NULL DEFAULT '',
    active_location BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Isso precisa de um trigger para ser atualizado automaticamente
    active BOOLEAN DEFAULT TRUE,
    CONSTRAINT unique_user_cnpj UNIQUE (cnpj),
    CONSTRAINT unique_user_cnh UNIQUE (cnh)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tb_users;
-- +goose StatementEnd
