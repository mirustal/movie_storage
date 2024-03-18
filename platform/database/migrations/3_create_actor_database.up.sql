CREATE TABLE IF NOT EXISTS actors (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    gender VARCHAR(50),
    birth_date DATE NOT NULL
);
