CREATE DATABASE wb_intern_1;
CREATE USER dev1 WITH PASSWORD '1a2s3d4f5g';
GRANT CREATE ON SCHEMA public TO dev1;
CREATE TABLE items (
    chrt_id bigserial not null primary key,
    track_number varchar not null,
    price numeric not null,
    rid varchar not null unique,
    name varchar not null,
    sale int not null,
    size varchar not null,
    total_price int not null,
    nm_id bigserial not null unique,
    brand varchar not null,
    status int not null
);
CREATE TABLE payment (
    transaction varchar not null primary key,
    request_id varchar not null,
    currency varchar not null,
    provider varchar not null,
    amount int not null,
    payment_dt bigserial not null unique,
    bank varchar not null,
    deliver_cost int not null,
    goods_total int not null,
    custom_fee int not null
);
CREATE TABLE orders (
    id varchar not null primary key,
    info jsonb not null unique
);