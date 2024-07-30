CREATE TABLE IF NOT EXISTS answers(
    id SERIAL PRIMARY KEY,
    answer TEXT,
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)
