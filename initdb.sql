DROP DATABASE ecommerce;
CREATE DATABASE ecommerce;
use ecommerce;

CREATE TABLE IF NOT EXISTS cart_products (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    user_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(128) NOT NULL,
    password VARCHAR(256) NOT NULL,
    credit INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(128) NOT NULL,
    price INTEGER NOT NULL,
    stock INTEGER NOT NULL
);

INSERT INTO products (name, price, stock)
VALUES 
    ('ipad pro 1', 100, 10),
    ('ipad pro 2', 200, 10),
    ('ipad pro 3', 300, 10),
    ('ipad pro 4', 400, 10),
    ('ipad pro 5', 500, 10),
    ('ipad pro 6', 600, 10),
    ('ipad pro 7', 700, 10),
    ('ipad pro 8', 800, 10);
