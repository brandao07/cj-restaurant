DROP TRIGGER IF EXISTS set_updated_at_customers ON customers;

CREATE TRIGGER set_updated_at_customers
BEFORE UPDATE ON customers
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS set_updated_at_restaurant_tables ON restaurant_tables;

CREATE TRIGGER set_updated_at_restaurant_tables
BEFORE UPDATE ON restaurant_tables
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();


DROP TRIGGER IF EXISTS trg_set_table_available ON customers;
-- Trigger on customers table
CREATE TRIGGER trg_set_table_available
AFTER UPDATE OF status ON customers
FOR EACH ROW
EXECUTE FUNCTION set_table_available_when_customer_done();