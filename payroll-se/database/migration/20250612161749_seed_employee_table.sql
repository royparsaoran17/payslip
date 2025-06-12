-- +goose Up
INSERT INTO employees (id, user_id, salary, created_at, updated_at, deleted_at)
VALUES ('105b5ac4-53c4-11ee-8c99-0242ac120002', 'a05b5ac4-53c4-11ee-8c99-0242ac120001', 1000000,  '2023-10-09 19:18:05.000000',
        '2023-10-09 19:18:05.000000', null);

-- +goose Down
DELETE
FROM employees
WHERE id = '105b5ac4-53c4-11ee-8c99-0242ac120002';

