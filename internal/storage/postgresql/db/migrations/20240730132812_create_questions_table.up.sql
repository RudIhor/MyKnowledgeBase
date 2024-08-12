CREATE TABLE IF NOT EXISTS questions(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    user_id INTEGER REFERENCES users(id),
    is_answered BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)
