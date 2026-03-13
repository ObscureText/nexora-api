CREATE TABLE user_roles (
    id INT PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

INSERT INTO user_roles (id, name) VALUES
(1,'COMPANY_ADMIN'),
(2,'INDIVIDUAL');
