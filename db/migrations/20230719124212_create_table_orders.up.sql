CREATE TABLE orders (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    order_number VARCHAR(255) NOT NULL,
    merchat_id VARCHAR(255) NOT NULL,
    product_id VARCHAR(255) NOT NULL,
    qty INT(10) NOT NULL,
    price float(50) NOT NULL,
    discount float(50) NOT NULL,
    status char(10),
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME

) ENGINE = InnoDB;