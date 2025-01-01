-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO tb_motorcycles (year, model, plate)
VALUES
(2004, 'Kawasaki Ninja', '123123'),
(2024, 'Honda', '234234'),
(1998, 'Harley', '345345');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
