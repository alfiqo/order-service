CREATE TABLE customers (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    dob DATE DEFAULT NULL,
    phone VARCHAR(15),
    address TEXT,
    updated_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    deleted_at DATETIME,

    UNIQUE(id)
) ENGINE = InnoDB;