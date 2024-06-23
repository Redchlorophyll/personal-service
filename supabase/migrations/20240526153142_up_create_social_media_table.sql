CREATE TABLE IF NOT EXISTS social_media_table (
    id SERIAL PRIMARY KEY,
    social_media VARCHAR(50) NOT NULL UNIQUE,
    url VARCHAR(255),
    account_id INTEGER,
    CONSTRAINT fk_account
        FOREIGN KEY (account_id)
        REFERENCES account_table (id)
);