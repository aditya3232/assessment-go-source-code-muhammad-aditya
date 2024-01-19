create table invoices
(
    id              varchar(255) not null,
    invoice_number  varchar(255) not null,
    customer_id     varchar(255) not null,
    subject         varchar(255) not null,
    issued_date     bigint       not null,
    due_date        bigint       not null,
    total_item      bigint       not null,
    sub_total       bigint       not null,
    grand_total     bigint       not null,
    status          varchar(255) not null,
    created_at      bigint       not null,
    updated_at      bigint       not null,
    primary key (id),
    foreign key fk_invoices_customer_id (customer_id) references customers (id)
) engine = innodb;