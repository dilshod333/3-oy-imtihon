create table if not exists products(
    product_id  serial primary key,
    name text, 
    type varchar(64),
    quantity int,
    description text,
    price float,
    created_at text
);
