create table users(
    user_id serial primary key,
    name varchar(64),
    age int,
    email varchar(64),
    password varchar(64),
    created_at text
);