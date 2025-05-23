DROP FUNCTION IF EXISTS update_updated_at_column() CASCADE;

CREATE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Function to update restaurant_table status
DROP FUNCTION IF EXISTS set_table_available_when_customer_done() CASCADE;
CREATE FUNCTION set_table_available_when_customer_done()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status IN ('NEEDS_PAYMENT', 'PAID') AND (OLD.status IS DISTINCT FROM NEW.status) THEN
        UPDATE restaurant_tables
        SET status = 'AVAILABLE'
        WHERE id = NEW.restaurant_table_id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;