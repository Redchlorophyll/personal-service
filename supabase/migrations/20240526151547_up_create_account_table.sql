CREATE TABLE IF NOT EXISTS account_table (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    profile_img_url VARCHAR(255),
    user_name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    full_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    salt CHAR(32) NOT NULL,
    session_token VARCHAR(255),
    last_logged_in TIMESTAMP,
    deleted_at TIMESTAMP
);