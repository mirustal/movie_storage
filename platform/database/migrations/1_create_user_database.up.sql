    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        refresh_token VARCHAR,
        role VARCHAR(255) NOT NULL
    );