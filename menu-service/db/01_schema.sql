CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
	created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL,
    name TEXT NOT NULL UNIQUE, 
    type TEXT NOT NULL CHECK (
        type IN ('STARTER', 'MAIN', 'DESSERT', 'DRINK') 
    )
);

CREATE TABLE IF NOT EXISTS menus (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL,
    starter_id INT REFERENCES items(id), 
    main_id INT REFERENCES items(id), 
    dessert_id INT REFERENCES items(id), 
    drink_id INT REFERENCES items(id), 
    price REAL
);