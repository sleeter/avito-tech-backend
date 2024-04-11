package storage

import (
	"avito-tech-backend/internal/pkg/pgdb"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	URL string `yaml:"url" env-required:"true"`
}

type Storage struct {
	Config   Config
	Database *pgdb.Database

	Banners BannerMapper
}

func NewStorage(ctx context.Context, cfg Config) (*Storage, error) {
	var err error
	storage := &Storage{
		Config: cfg,
	}
	pool, err := pgxpool.Connect(ctx, cfg.URL)
	if err != nil {
		return nil, err
	}
	storage.Database = pgdb.NewDatabase(pool)
	storage.Banners = BannerMapper{Storage: storage}
	return storage, nil
}
