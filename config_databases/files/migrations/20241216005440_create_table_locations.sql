-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS tb_locations (
    id INTEGER, 
    user_id INTEGER NOT NULL,
    motorcycle_id INTEGER NOT NULL,
    price REAL DEFAULT 0,
    days INTEGER DEFAULT 0,
    location_day TIMESTAMP NOT NULL,
    devolution_day TIMESTAMP NOT NULL,
    fine REAL DEFAULT 0,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_location_id PRIMARY KEY (id AUTOINCREMENT),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES tb_users(id),
    CONSTRAINT fk_motorcycle FOREIGN KEY(motorcycle_id) REFERENCES tb_motorcycles(id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
