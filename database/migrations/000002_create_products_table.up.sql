CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL DEFAULT 'Unknown Product',
			price REAL NOT NULL DEFAULT 0.0,
			description TEXT NOT NULL DEFAULT 'No description',
			image_url TEXT NOT NULL DEFAULT 'https://example.com/default.jpg',
			quantity INTEGER NOT NULL DEFAULT 0
		);