-- +goose Up
create table if not exists reimbursements
(
    id          uuid
    constraint reimbursements_pk
    primary key,
    employee_id         varchar   not null,
    reimbursement_date  date not null,
    amount              numeric(15,2) not null,
    description         text,
    status              varchar,
    created_at          timestamp not null,
    updated_at          timestamp not null,
    deleted_at          timestamp,
    created_by          varchar,
    updated_by          varchar,
    deleted_by          varchar
);

-- +goose Down
drop table if exists reimbursements;
