-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO tb_locations ('user_id', 'motorcycle_id', 'price', 'days', 'location_day', 'devolution_day', 'fine')
VALUES
    (1, 1, 100, 5, '2024-07-22', '2024-07-27', 0),
    (2, 2, 100, 5, '2024-07-23', '2024-07-29', 10),
    (3, 3, 100, 7, '2024-07-23', '2024-07-29', 0);

INSERT INTO tb_locations ('user_id', 'motorcycle_id', 'price', 'days', 'location_day', 'devolution_day', 'fine', 'active')
VALUES
    (1, 2, 100, 19, '2024-01-01', '2024-01-20', 200, FALSE);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
