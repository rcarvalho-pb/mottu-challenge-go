-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tb_user_motorcycle(
    user_id INTEGER NOT NULL,
    motorcycle_id INTEGER NOT NULL,
    CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES tb_users (id) ON DELETE CASCADE,
    CONSTRAINT motorcycle_fk FOREIGN KEY (motorcycle_id) REFERENCES tb_motorcycles (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tb_user_motorcycle;
-- +goose StatementEnd
