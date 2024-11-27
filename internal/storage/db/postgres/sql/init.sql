CREATE TABLE IF NOT EXISTS users (
                       id SERIAL PRIMARY KEY,
                       last_name VARCHAR(100) NOT NULL,
                       first_name VARCHAR(100) NOT NULL,
                       gender VARCHAR(100) NOT NULL,
                       birth_date VARCHAR(100) NOT NULL,
                       phone VARCHAR(100) NOT NULL UNIQUE
);