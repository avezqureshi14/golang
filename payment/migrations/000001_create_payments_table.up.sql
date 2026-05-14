CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    amount INT NOT NULL,
    status TEXT NOT NULL
);