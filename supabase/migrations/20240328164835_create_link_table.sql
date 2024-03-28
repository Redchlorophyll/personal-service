CREATE TABLE IF NOT EXISTS link (
    id SERIAL PRIMARY KEY,
    img_url VARCHAR(255),
    description VARCHAR(255) NOT NULL,
    url_anchor VARCHAR(255) NOT NULL,
    deleted_at TIMESTAMP
);