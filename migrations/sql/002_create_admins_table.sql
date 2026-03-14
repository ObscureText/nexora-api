CREATE TABLE admins (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    mobile_no TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);