DROP TRIGGER IF EXISTS set_updated_at_restaurant_orders ON restaurant_orders;

CREATE TRIGGER set_updated_at_restaurant_orders
BEFORE UPDATE ON restaurant_orders
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();