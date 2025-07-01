CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    birth_date DATE NOT NULL,
    gender TEXT NOT NULL,
    interests TEXT,
    city TEXT,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);
