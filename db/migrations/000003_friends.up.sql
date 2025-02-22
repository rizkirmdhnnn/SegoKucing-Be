CREATE TABLE friends (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    friend_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (friend_id) REFERENCES users(id)
);

CREATE INDEX idx_friends_user_id ON friends(user_id);

CREATE INDEX idx_friends_friend_id ON friends(friend_id);