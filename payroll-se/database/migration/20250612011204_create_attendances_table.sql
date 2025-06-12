-- +goose Up
create table if not exists attendances
(
    id          uuid
    constraint attendances_pk
    primary key,
    employee_id         varchar   not null,
    attendance_date     date   not null,
    created_at          timestamp not null,
    updated_at          timestamp not null,
    deleted_at          timestamp,
    created_by          varchar,
    updated_by          varchar,
    deleted_by          varchar,
    constraint no_weekends check (extract(dow from attendance_date) not in (0,6))
);

-- +goose Down
drop table if exists attendances;
