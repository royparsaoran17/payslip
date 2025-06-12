-- +goose Up
create table if not exists payroll_periods
(
    id          uuid
    constraint payroll_periods_pk
    primary key,
    start_date      date   not null,
    end_date        date   not null,
    is_processed    boolean default false,
    created_at      timestamp not null,
    updated_at      timestamp not null,
    deleted_at      timestamp,
    created_by      varchar,
    updated_by      varchar,
    deleted_by      varchar
);

-- +goose Down
drop table if exists payroll_periods;
