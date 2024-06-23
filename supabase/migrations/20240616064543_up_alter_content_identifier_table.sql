-- step 1: alter table, add account_id.
ALTER TABLE content_identifier
ADD COLUMN account_url_slug_id INTEGER;

-- step 2: add foreign key constraint to account_id.
ALTER TABLE content_identifier
ADD CONSTRAINT fk_account
FOREIGN KEY (account_url_slug_id)
REFERENCES account_url_slug_table (id);


-- step 3: indexing the foreign_key.
CREATE INDEX idx_slug_url_id ON content_identifier (account_url_slug_id);