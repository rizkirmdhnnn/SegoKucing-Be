CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    tag TEXT NOT NULL
);

CREATE INDEX idx_tags_tag ON tags(tag);