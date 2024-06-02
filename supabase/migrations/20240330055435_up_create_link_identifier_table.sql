CREATE TABLE IF NOT EXISTS link_identifier (
    link_id INT,
    type_id INT,
    PRIMARY KEY (link_id, type_id),
    FOREIGN KEY () REFERENCES link(id),
link_id    FOREIGN KEY (type_id) REFERENCES content_identifier(id)
);