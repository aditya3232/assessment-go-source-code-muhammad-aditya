create table invoice_items
(
    id              varchar(255) not null,
    invoice_id      varchar(255) not null,
    item_id         varchar(255) not null,
    item_quantity   int          not null,
    amount          bigint       not null,
    created_at      bigint       not null,
    updated_at      bigint       not null,
    primary key (id),
    foreign key fk_invoice_items_invoice_id (invoice_id) references invoices (id),
    foreign key fk_invoice_items_item_id (item_id) references items (id)
) engine = innodb;