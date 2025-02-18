CREATE TABLE post_tags (
    id SERIAL PRIMARY KEY,
    posts_id INT NOT NULL,
    tags_id INT NOT NULL,
    FOREIGN KEY (posts_id) REFERENCES posts (id),
    FOREIGN KEY (tags_id) REFERENCES tags (id)
);

CREATE INDEX idx_post_tags_posts_id ON post_tags(posts_id);