INSERT INTO service_providers (id, name, email, mobile_number, address) 
VALUES ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Desert Business Solutions', 'contact@dbs.ae', '0500000000', 'Dubai, UAE');

INSERT INTO users (id, service_provider_id, name, email, password_hash, role_id, mobile_number) 
VALUES 
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Provider Admin', 'admin@dbs.ae', '$2a$10$EruQUo5d1xDJp1XX1gB6nuFc/KkMkw9o3WCRWpRL3bxpoGB7DEkpG', 1, '0501111111'),
    ('cccccccc-cccc-cccc-cccc-cccccccccccc', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Establishment Owner', 'owner@company.com', '$2a$10$EruQUo5d1xDJp1XX1gB6nuFc/KkMkw9o3WCRWpRL3bxpoGB7DEkpG', 2, '0502222222'),
    ('dddddddd-dddd-dddd-dddd-dddddddddddd', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Individual User', 'individual@email.com', '$2a$10$EruQUo5d1xDJp1XX1gB6nuFc/KkMkw9o3WCRWpRL3bxpoGB7DEkpG', 3, '0503333333');
