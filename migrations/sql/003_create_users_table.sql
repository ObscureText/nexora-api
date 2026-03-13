CREATE TABLE users (
    id UUID PRIMARY KEY,
    admin_id UUID NOT NULL REFERENCES admins(id),
    role_id INT NOT NULL REFERENCES user_roles(id),
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    mobile_no TEXT,
    created_at TIMESTAMP NOT NULL
);
