CREATE TABLE IF NOT EXISTS "Customer" (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    document VARCHAR(50) NOT NULL,
    creation_date TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "Vehicle"(
    id VARCHAR(255) PRIMARY KEY,
    make VARCHAR(100) NOT NULL,
    model VARCHAR(100) NOT NULL,
    year INT NOT NULL,
    paint VARCHAR(50) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    available BOOLEAN NOT NULL,
    creation_date TIMESTAMP NOT NULL,
    modified_date TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "Sale" (
    id VARCHAR(255) PRIMARY KEY,
    data TIMESTAMP NOT NULL,
    vehicle_id VARCHAR(255) NOT NULL,
    customer_id VARCHAR(255) NOT NULL,
    payment_type VARCHAR(50) NOT NULL,
    FOREIGN KEY (vehicle_id) REFERENCES "Vehicle" (id),
    FOREIGN KEY (customer_id) REFERENCES "Customer" (id)
);
