CREATE TABLE users (
    id UUID PRIMARY KEY,
    service_provider_id UUID NOT NULL REFERENCES service_providers(id),
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role_id INT NOT NULL REFERENCES user_roles(id),
    mobile_number TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
