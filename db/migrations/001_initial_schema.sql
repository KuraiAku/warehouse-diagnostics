CREATE TABLE warehouses (
    warehouse_id INT IDENTITY(1,1) PRIMARY KEY,
    warehouse_code VARCHAR(20) NOT NULL UNIQUE,
    warehouse_name VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    state_code CHAR(2) NOT NULL,
    is_active BIT NOT NULL DEFAULT 1
);


CREATE TABLE products (
    product_id INT IDENTITY(1,1) PRIMARY KEY,
    sku VARCHAR(50) NOT NULL UNIQUE,
    product_name VARCHAR(150) NOT NULL,
    category VARCHAR(100) NULL,
    unit_cost DECIMAL(10,2) NOT NULL
);

CREATE TABLE inventory (
    inventory_id INT IDENTITY(1,1) PRIMARY KEY,
    warehouse_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    location_code VARCHAR(50) NULL,

    CONSTRAINT FK_inventory_warehouse
        FOREIGN KEY (warehouse_id)
        REFERENCES warehouses(warehouse_id),

    CONSTRAINT FK_inventory_product
        FOREIGN KEY (product_id)
        REFERENCES products(product_id),

    CONSTRAINT UQ_inventory_warehouse_product
        UNIQUE (warehouse_id, product_id)
);

CREATE TABLE orders (
    order_id BIGINT IDENTITY(1,1) PRIMARY KEY,
    warehouse_id INT NOT NULL,
    status VARCHAR(30) NOT NULL,
    total_amount DECIMAL(12,2) NOT NULL DEFAULT 0,
    created_at DATETIME2 NOT NULL DEFAULT SYSUTCDATETIME(),

    CONSTRAINT FK_orders_warehouse
        FOREIGN KEY (warehouse_id)
        REFERENCES warehouses(warehouse_id)
);

CREATE TABLE order_items (
    order_item_id BIGINT IDENTITY(1,1) PRIMARY KEY,
    order_id BIGINT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,

    CONSTRAINT FK_order_items_order
        FOREIGN KEY (order_id)
        REFERENCES orders(order_id),

    CONSTRAINT FK_order_items_product
        FOREIGN KEY (product_id)
        REFERENCES products(product_id)
);


CREATE TABLE inventory_movements (
    movement_id BIGINT IDENTITY(1,1) PRIMARY KEY,
    warehouse_id INT NOT NULL,
    product_id INT NOT NULL,
    order_id BIGINT NULL,
    movement_type VARCHAR(30) NOT NULL,
    quantity_change INT NOT NULL,
    created_at DATETIME2 NOT NULL DEFAULT SYSUTCDATETIME(),

    CONSTRAINT FK_inventory_movements_warehouse
        FOREIGN KEY (warehouse_id)
        REFERENCES warehouses(warehouse_id),

    CONSTRAINT FK_inventory_movements_product
        FOREIGN KEY (product_id)
        REFERENCES products(product_id),

    CONSTRAINT FK_inventory_movements_order
        FOREIGN KEY (order_id)
        REFERENCES orders(order_id)
);
