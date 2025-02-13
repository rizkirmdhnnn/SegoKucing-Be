CREATE TABLE post_tags (
    id SERIAL PRIMARY KEY,
    post_id INT NOT NULL,
    tag_id INT NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts (id),
    FOREIGN KEY (tag_id) REFERENCES tags (id)
);