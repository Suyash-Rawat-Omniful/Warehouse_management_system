CREATE TABLE IF NOT EXISTS inventories (
    id SERIAL PRIMARY KEY,
    sku_id INT NOT NULL,
    hub_id INT NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
