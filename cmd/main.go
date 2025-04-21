package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hferr/hw-rest-api/config"
	"github.com/hferr/hw-rest-api/internal/app"
	"github.com/hferr/hw-rest-api/internal/httpjson"
	"github.com/hferr/hw-rest-api/internal/repository/psql"
	"github.com/hferr/hw-rest-api/migrations"
)

const fmtDBConnString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
	ctx := context.Background()
	cfg := config.New()

	db, err := initPostgresDb(ctx, cfg.DB)
	if err != nil {
		log.Fatalf("failed to setup database: %v", err)
	}
	defer db.Close()

	repo := psql.NewRepo(db.Db)

	agentSvc := app.NewAgentService(repo)
	customerSvc := app.NewCustomerService(repo)

	handler := httpjson.NewHandler(
		agentSvc,
		customerSvc,
	)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      handler.NewRouter(),
		ReadTimeout:  cfg.Server.TimeoutRead,
		WriteTimeout: cfg.Server.TimeoutWrite,
		IdleTimeout:  cfg.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

func initPostgresDb(ctx context.Context, cfg config.CfgDB) (*psql.Postgres, error) {
	dbConnString := fmt.Sprintf(
		fmtDBConnString,
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
	)

	psql, err := psql.NewPostgresDb(ctx, dbConnString)
	if err != nil {
		return nil, err
	}

	if err := migrations.MaybeApplyMigrations(psql.Db); err != nil {
		return nil, err
	}

	return psql, nil
}
