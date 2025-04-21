-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS applications(
    id uuid PRIMARY KEY,
    customer_id uuid REFERENCES customers(id) NOT NULL,
    agent_id uuid REFERENCES agents(id) NOT NULL,
    purchasing_address VARCHAR(255) NOT NULL,
    application_approved BOOLEAN NOT NULL DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS applications;
-- +goose StatementEnd
