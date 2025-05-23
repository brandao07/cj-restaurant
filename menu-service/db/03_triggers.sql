DROP TRIGGER IF EXISTS set_updated_at_menus ON menus;

CREATE TRIGGER set_updated_at_menus
BEFORE UPDATE ON menus
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

DROP TRIGGER IF EXISTS set_updated_at_items ON items;

CREATE TRIGGER set_updated_at_items
BEFORE UPDATE ON items
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();