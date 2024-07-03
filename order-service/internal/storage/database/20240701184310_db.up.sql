create table if not exists orders(
    order_id serial primary key,
    user_id int 
    product_id int 
    name text,
    price float
    total_price float,
    order_time text
);