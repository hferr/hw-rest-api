package psql

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/hferr/hw-rest-api/internal/app"
)

const FindCustomerByIDQuery = `SELECT * FROM customers WHERE id = $1`

func (r *Repo) FindCustomerByID(ctx context.Context, ID uuid.UUID) (app.Customer, error) {
	row := r.db.QueryRowContext(ctx, FindCustomerByIDQuery, ID)

	customer, err := mapOutCustomer(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return app.Customer{}, app.ErrCustomerNotFound
		}
		return app.Customer{}, app.ErrInternal
	}

	return customer, nil
}

const CheckCustomerAgentConnectionExistsQuery = `
	SELECT EXISTS (
		SELECT a.* FROM agents a
		LEFT JOIN applications app ON app.agent_id = a.id
		WHERE app.customer_id = $1 AND app.agent_id = $2
	)
`

func (r *Repo) IsCustomerConnectedToAgent(ctx context.Context, customerID, agentID uuid.UUID) (bool, error) {
	row := r.db.QueryRowContext(ctx, CheckCustomerAgentConnectionExistsQuery, customerID, agentID)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, app.ErrInternal
	}

	return exists, nil
}

func mapOutCustomer(row *sql.Row) (app.Customer, error) {
	var customer app.Customer
	err := row.Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.PhoneNumber,
		&customer.CurrentAddress,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)
	if err != nil {
		return app.Customer{}, err
	}

	return customer, nil
}
