create table if not exists inv_sell_per_item
(
    id           integer not null auto_increment,
    commodity_id integer not null,
    sell_price   integer not null,
    created_at   timestamp default now(),
    updated_at   timestamp    null,
    primary key (id),
    foreign key (commodity_id) references inv_commodity (id)
);