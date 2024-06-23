-- step 1: create account table
CREATE TABLE IF NOT EXISTS account_table (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    profile_img_url VARCHAR(255),
    user_name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    full_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    session_token VARCHAR(255) UNIQUE,
    session_token_expired_at TIMESTAMP,
    last_logged_in TIMESTAMP,
    last_logged_out TIMESTAMP,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    refresh_token VARCHAR(255) UNIQUE,
    refresh_token_expired_at VARCHAR(255),
    updated_at VARCHAR(255)
);

-- step 2: Index on email column & user_name
CREATE INDEX idx_email ON account_table (email);
CREATE INDEX idx_user_name ON account_table (user_name);
