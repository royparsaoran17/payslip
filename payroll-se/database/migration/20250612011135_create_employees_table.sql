-- +goose Up
create table if not exists employees
(
    id          uuid
    constraint employees_pk
    primary key,
    user_id     varchar   not null,
    salary      numeric(15,2)   not null,
    created_at  timestamp not null,
    updated_at  timestamp not null,
    deleted_at  timestamp,
    created_by  varchar,
    updated_by  varchar,
    deleted_by  varchar
    );

create index employees_user_id_index on employees (user_id);

-- +goose Down
drop table if exists employees;
