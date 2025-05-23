CREATE TABLE restaurant_tables (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    status VARCHAR(32) NOT NULL DEFAULT 'AVAILABLE' CHECK (
        status IN ('AVAILABLE', 'UNAVAILABLE')
    )
);

CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    status VARCHAR(32) NOT NULL DEFAULT 'WAITING' CHECK (
        status IN ('WAITING', 'EATING', 'NEEDS_PAYMENT', 'PAID')
    ),
    restaurant_table_id INTEGER NOT NULL REFERENCES restaurant_tables(id),
    order_id INTEGER
);




