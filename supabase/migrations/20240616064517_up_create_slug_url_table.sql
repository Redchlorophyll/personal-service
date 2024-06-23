CREATE TABLE IF NOT EXISTS account_url_slug_table (
    id SERIAL PRIMARY KEY,
    url_slug VARCHAR(50) NOT NULL UNIQUE,
    account_id INTEGER,
    CONSTRAINT fk_account
        FOREIGN KEY (account_id)
        REFERENCES account_table (id)
);