package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hferr/hw-rest-api/config"
	"github.com/hferr/hw-rest-api/internal/httpjson"
)

func main() {
	cfg := config.New()

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
