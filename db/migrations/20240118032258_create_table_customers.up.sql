create table customers
(
    id              varchar(255) not null,
    national_id     varchar(255) not null,
    name            varchar(255) not null,
    detail_address  text not null,
    created_at      bigint       not null,
    updated_at      bigint       not null,
    primary key (id)
) engine = InnoDB;