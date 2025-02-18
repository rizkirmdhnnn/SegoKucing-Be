CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT DEFAULT NULL UNIQUE,
    phone TEXT DEFAULT NULL UNIQUE,
    password TEXT NOT NULL,
    image_url TEXT DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);

CREATE INDEX idx_users_phone ON users(phone);