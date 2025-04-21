package psql

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/hferr/hw-rest-api/internal/app"
)

const FindAgentByIDQuery = `SELECT * FROM agents WHERE id = $1`

func (r *Repo) FindAgentByID(ctx context.Context, ID uuid.UUID) (app.Agent, error) {
	row := r.db.QueryRowContext(ctx, FindAgentByIDQuery, ID)

	out, err := mapOutAgentFromRow(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return app.Agent{}, app.ErrAgentNotFound
		}
		return app.Agent{}, app.ErrInternal
	}

	return out, nil
}

const UpsertAgentQuery = `
	INSERT INTO agents (
		id, name, email, phone_number, location, created_at, updated_at
	)
	VALUES (
		$1, $2, $3, $4, $5, $6, $7
	) ON CONFLICT (id) DO UPDATE SET
		name = $2,
		email = $3,
		phone_number = $4,
		location = $5,
		updated_at = NOW()
	returning *
`

func (r *Repo) UpsertAgent(ctx context.Context, agent app.Agent) (app.Agent, error) {
	row := r.db.QueryRowContext(
		ctx,
		UpsertAgentQuery,
		agent.ID,
		agent.Name,
		agent.Email,
		agent.PhoneNumber,
		agent.Location,
		agent.CreatedAt,
		agent.UpdatedAt,
	)

	out, err := mapOutAgentFromRow(row)
	if err != nil {
		return app.Agent{}, err
	}

	return out, nil
}

func mapOutAgentFromRow(row *sql.Row) (app.Agent, error) {
	var agent app.Agent
	err := row.Scan(
		&agent.ID,
		&agent.Name,
		&agent.Email,
		&agent.PhoneNumber,
		&agent.Location,
		&agent.CreatedAt,
		&agent.UpdatedAt,
	)
	if err != nil {
		return app.Agent{}, err
	}

	return agent, nil
}
