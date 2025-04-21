package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/hferr/hw-rest-api/config"
	"github.com/hferr/hw-rest-api/internal/httpjson"
	"github.com/hferr/hw-rest-api/internal/repository/psql"
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

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      httpjson.NewHandler().NewRouter(),
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

	return psql, nil
}
