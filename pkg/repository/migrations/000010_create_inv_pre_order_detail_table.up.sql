create table if not exists inv_pre_order_detail
(
    id      integer not null auto_increment,
    po_item text    null,
    primary key (id)
);