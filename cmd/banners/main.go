package main

import (
	"avito-tech-backend/internal/core"
	http_server "avito-tech-backend/internal/http-server"
	"avito-tech-backend/internal/pkg/config"
	"context"
	"database/sql"
	"errors"
	"github.com/avast/retry-go/v4"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"log"
	"log/slog"
	"os"
	"time"
)

func main() {
	ctx := context.Background()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	loader := config.PrepareLoader(config.WithConfigPath("./config.yaml"))

	cfg, err := core.ParseConfig(loader)
	if err != nil {
		log.Fatalf("Failed to parse config: %s", err)
	}

	err = retry.Do(func() error {
		return UpMigrations(cfg)
	}, retry.Attempts(4), retry.Delay(2*time.Second))
	if err != nil {
		log.Fatalf(err.Error())
	}

	repository, err := core.NewRepository(ctx, cfg)
	if err != nil {
		log.Fatalf("Init repository: %s", err)
	}
	app := http_server.New(repository)

	if err := app.Start(ctx); err != nil {
		log.Fatalf(err.Error())
	}

}

func UpMigrations(cfg *core.Config) error {
	db, err := sql.Open("pgx", cfg.Storage.URL)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil

}
