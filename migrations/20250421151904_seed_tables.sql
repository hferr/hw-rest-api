-- +goose Up
-- +goose StatementBegin
INSERT INTO agents (
    id, name, email, phone_number, location, created_at, updated_at
) VALUES
    (gen_random_uuid(), 'Michael Agent', 'm.agent@sample.com', '111-123-1234', 'Austin, TX', NOW(), NOW()),
    (gen_random_uuid(), 'Blake Agent', 'b.agent@sample.com', '999-999-9999', 'Atlanta, GA', NOW(), NOW()),
    (gen_random_uuid(), 'Robbie Agent', 'robbie.agent@sample.com', '111-111-1111', 'Denver, CO', NOW(), NOW());

INSERT INTO customers (
    id, name, email, phone_number, current_address, created_at, updated_at
) VALUES
    (gen_random_uuid(), 'Cynthia Hart', 'c.hart@sample.com', '111-123-1234', '908 Test St. Austin, TX 78749', NOW(), NOW()),
    (gen_random_uuid(), 'Russell Jackson', 'r.jackson@sample.com', '999-999-9999', '123 Test St. Atlanta, GA 30301', NOW(), NOW()),
    (gen_random_uuid(), 'Michelle Jones', 'michelle.jones@sample.com', '111-111-1111', '413 Testing St. Austin, TX 78703', NOW(), NOW());

-- seed applications linking agents with customers
INSERT INTO applications (id, customer_id, agent_id, purchasing_address, application_approved)
SELECT
	gen_random_uuid(), c.id, a.id, '111 Purchase St. Austin, TX 78739', false
FROM customers c, agents a
WHERE c.email='c.hart@sample.com' AND a.email='m.agent@sample.com'
UNION ALL
SELECT
	gen_random_uuid(), c.id, a.id,'355 Purchase St. Athens, GA 30304',true
FROM customers c, agents a
WHERE c.email='r.jackson@sample.com' AND a.email='b.agent@sample.com'
UNION ALL
SELECT
	gen_random_uuid(), c.id, a.id,'123 Purchasing St. Denver, CO 80014',true
FROM customers c, agents a
WHERE c.email='michelle.jones@sample.com' AND a.email='robbie.agent@sample.com';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM agents;
DELETE FROM customers;
DELETE FROM applications;
-- +goose StatementEnd
