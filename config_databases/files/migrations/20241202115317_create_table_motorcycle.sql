-- +goose Up
-- +goose StatementBegin
/*CREATE TABLE IF NOT EXISTS tb_motorcycles(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    year INT NOT NULL,
    model TEXT NOT NULL,
    plate TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_located BOOLEAN DEFAULT FALSE,
    active BOOLEAN DEFAULT TRUE,
    CONSTRAINT unique_plate UNIQUE (plate),
);
*/

CREATE TABLE IF NOT EXISTS tb_motorcycles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    year INTEGER NOT NULL,              -- Substituído INT por INTEGER
    model TEXT NOT NULL,
    plate TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Corrigido 'DEFAUL' para 'DEFAULT'
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Corrigido 'DEFAUL' para 'DEFAULT'
    is_located BOOLEAN DEFAULT FALSE,
    active BOOLEAN DEFAULT TRUE,
    CONSTRAINT unique_plate UNIQUE (plate)  -- Removida vírgula extra
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tb_motorcycles;
-- +goose StatementEnd
