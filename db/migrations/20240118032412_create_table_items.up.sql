create table items
(
    id                  varchar(255) not null,
    item_name           varchar(255) not null,
    type                varchar(255) not null,
    item_price          bigint       not null,
    created_at          bigint       not null,
    updated_at          bigint       not null,
    primary key (id)
) engine = InnoDB;