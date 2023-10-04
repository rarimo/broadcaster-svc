-- +migrate Up

create table transactions
(
    id            bigserial primary key,
    data          bytea   not null
);

-- +migrate Down

drop table transactions;