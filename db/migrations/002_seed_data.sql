INSERT INTO warehouses (
    warehouse_code,
    warehouse_name,
    city,
    state_code
)
VALUES
    ('ONT-01', 'Ontario Distribution Center', 'Ontario', 'CA'),
    ('NJ-01', 'New Jersey Distribution Center', 'Secaucus', 'NJ'),
    ('DAL-01', 'Dallas Distribution Center', 'Dallas', 'TX');



INSERT INTO products (
    sku,
    product_name,
    category,
    unit_cost
)
VALUES
    ('BOX-1001', 'Large Shipping Box', 'Packaging', 2.50),
    ('PAL-2001', 'Standard Wood Pallet', 'Warehouse', 18.00),
    ('WRAP-3001', 'Stretch Wrap Roll', 'Packaging', 7.25),
    ('TAPE-4001', 'Packing Tape', 'Packaging', 3.75),
    ('LBL-5001', 'Shipping Labels', 'Supplies', 12.00);

INSERT INTO inventory (
    warehouse_id,
    product_id,
    quantity,
    location_code
)
VALUES
    (1, 1, 100, 'A-01-01'),
    (1, 2, 45,  'B-02-03'),
    (1, 3, 20,  'C-01-02');

INSERT INTO inventory (
    warehouse_id,
    product_id,
    quantity,
    location_code
)
VALUES
    (2, 1, 250, 'A-01-01'),
    (2, 3, 15,  'C-02-01'),
    (2, 4, 80,  'D-01-04'),
    (3, 1, 60,  'A-03-02'),
    (3, 2, 12,  'B-01-01'),
    (3, 5, 200, 'E-02-05');

INSERT INTO orders (
    warehouse_id,
    status,
    total_amount
)
VALUES
    (1, 'pending', 125.50),
    (1, 'allocated', 78.00),
    (2, 'pending', 210.25),
    (3, 'picked', 95.00),
    (3, 'pending', 42.75);


INSERT INTO order_items (
    order_id,
    product_id,
    quantity,
    unit_price
)
VALUES
    (1, 1, 10, 2.50),
    (1, 3, 5, 7.25),

    (2, 2, 2, 18.00),

    (3, 1, 20, 2.50),
    (3, 4, 6, 3.75),

    (4, 5, 4, 12.00),

    (5, 3, 3, 7.25);


INSERT INTO inventory_movements (
    warehouse_id,
    product_id,
    order_id,
    movement_type,
    quantity_change
)
VALUES
    (1, 1, NULL, 'receive', 100),
    (1, 2, NULL, 'receive', 45),
    (1, 3, NULL, 'receive', 20),

    (1, 1, 1, 'pick', -10),
    (1, 3, 1, 'pick', -5),

    (2, 1, NULL, 'receive', 250),
    (2, 1, 3, 'pick', -20),

    (3, 5, NULL, 'receive', 200),
    (3, 3, 5, 'pick', -3);
