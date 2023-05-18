CREATE TABLE IF NOT EXISTS Items (
    item_id      SERIAL PRIMARY KEY,
    item_code    VARCHAR(100),
    description  VARCHAR(100),
    quantity     INTEGER,
    order_id     INTEGER,
    created_at	 TIMESTAMP,
    updated_at	 TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES Orders (order_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Orders (
    order_id      SERIAL PRIMARY KEY,
    customer_name VARCHAR(100),
    created_at	 TIMESTAMP,
    updated_at	 TIMESTAMP
);