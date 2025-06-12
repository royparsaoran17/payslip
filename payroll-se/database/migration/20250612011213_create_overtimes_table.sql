-- +goose Up
create table if not exists overtimes
(
    id          uuid
    constraint overtimes_pk
    primary key,
    employee_id     varchar   not null,
    overtime_date   date   not null,
    hours           numeric(3,1) not null check (hours <= 3),
    created_at      timestamp not null,
    updated_at      timestamp not null,
    deleted_at      timestamp,
    created_by      varchar,
    updated_by      varchar,
    deleted_by      varchar
);

-- +goose Down
drop table if exists overtimes;
