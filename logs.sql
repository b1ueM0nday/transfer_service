create table logs
(
    id           serial
        constraint logs_pk
            primary key,
    date         timestamp not null,
    message_type text      not null,
    message      json      not null
);

alter table logs
    owner to postgres;

create unique index logs_id_uindex
    on logs (id);

