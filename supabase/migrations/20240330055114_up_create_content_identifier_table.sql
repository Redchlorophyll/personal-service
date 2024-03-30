CREATE TABLE IF NOT EXISTS content_identifier (
    id SERIAL PRIMARY KEY,
    identifier VARCHAR(50) NOT NULL,
    title VARCHAR(255) NOT NULL
);