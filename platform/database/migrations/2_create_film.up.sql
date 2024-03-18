CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description TEXT CHECK (length(description) <= 1000),
    release_date DATE NOT NULL,
    rating DECIMAL(2, 1) CHECK (rating >= 0 AND rating <= 10)
);