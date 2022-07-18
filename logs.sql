create table logs
(
    id           serial
        constraint logs_pk
            primary key,
    date         timestamp not null,
    op_type text      not null,
    message      json      not null
);

create table receipts
(
    id           serial
        constraint receipts_pk
            primary key,
    date    timestamp default CURRENT_TIMESTAMP                   not null,
    op_type text      default 'Undefined'::text,
    receipt json
);
