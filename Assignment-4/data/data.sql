CREATE DATABASE mystore;

USE mystore;

CREATE TABLE cart (
    ItemName     VARCHAR(50),
    ItemPrice    FLOAT(3),
	ItemType     VARCHAR(20),
	ItemQuantity INT
);

INSERT INTO cart (ItemName, ItemPrice, ItemType, ItemQuantity) 
VALUES
("Tshirt", 299.43, "raw", 2),
("Jeans", 1800.56, "manufactured", 1),
("Shirt", 1200, "imported", 3),
("Kurta", 800.24, "raw", 6);

