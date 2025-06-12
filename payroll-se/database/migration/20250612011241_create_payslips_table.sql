-- +goose Up
create table if not exists payslips
(
    id          uuid
    constraint payslips_pk
    primary key,
    employee_id         varchar   not null,
    payroll_period_id   varchar not null,
    base_salary         numeric(15,2) not null,
    prorated_salary     numeric(15,2) not null,
    overtime_pay        numeric(15,2) not null,
    reimbursement_total numeric(15,2) not null,
    take_home_pay       numeric(15,2) not null,
    created_at          timestamp not null,
    updated_at          timestamp not null,
    deleted_at          timestamp,
    created_by          varchar,
    updated_by          varchar,
    deleted_by          varchar
);

-- +goose Down
drop table if exists payslips;
