CREATE TABLE IF NOT EXISTS order_items (
			id SERIAL PRIMARY KEY,
			order_id INTEGER NOT NULL,
			product_id INTEGER NOT NULL,
			quantity INTEGER NOT NULL CHECK(quantity > 0),
			price REAL NOT NULL CHECK(price >= 0),
			FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
			FOREIGN KEY (product_id) REFERENCES products(id)
		);