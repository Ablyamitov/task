CREATE TABLE IF NOT EXISTS users (
                       id SERIAL PRIMARY KEY,
                       last_name VARCHAR(100) NOT NULL,
                       first_name VARCHAR(100) NOT NULL,
                       gender VARCHAR(100) NOT NULL,
                       birth_date VARCHAR(100) NOT NULL,
                       phone VARCHAR(100) NOT NULL UNIQUE,
                       role VARCHAR(100) NOT NULL
);

insert into users (last_name, first_name, gender, birth_date, phone, role)
VALUES ('admin', 'admin', 'male', '10-10-2000', '+79787678178', 'Role_Admin')