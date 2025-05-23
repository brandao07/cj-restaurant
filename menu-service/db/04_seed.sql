-- Starters
INSERT INTO items (name, type) VALUES ('Tomato Soup', 'STARTER');
INSERT INTO items (name, type) VALUES ('Bruschetta', 'STARTER');
INSERT INTO items (name, type) VALUES ('Caesar Salad', 'STARTER');

-- Mains
INSERT INTO items (name, type) VALUES ('Grilled Salmon', 'MAIN');
INSERT INTO items (name, type) VALUES ('Steak Frites', 'MAIN');
INSERT INTO items (name, type) VALUES ('Vegetable Lasagna', 'MAIN');

-- Desserts
INSERT INTO items (name, type) VALUES ('Chocolate Cake', 'DESSERT');
INSERT INTO items (name, type) VALUES ('Panna Cotta', 'DESSERT');
INSERT INTO items (name, type) VALUES ('Fruit Tart', 'DESSERT');

-- Drinks
INSERT INTO items (name, type) VALUES ('Red Wine', 'DRINK');
INSERT INTO items (name, type) VALUES ('Sparkling Water', 'DRINK');
INSERT INTO items (name, type) VALUES ('Espresso', 'DRINK');

-- Menu
INSERT INTO menus (starter_id, main_id, dessert_id, drink_id, price)
VALUES (
    (SELECT id FROM items WHERE type = 'STARTER' ORDER BY id LIMIT 1),
    (SELECT id FROM items WHERE type = 'MAIN' ORDER BY id LIMIT 1),
    (SELECT id FROM items WHERE type = 'DESSERT' ORDER BY id LIMIT 1),
    (SELECT id FROM items WHERE type = 'DRINK' ORDER BY id LIMIT 1),
    29.99
);