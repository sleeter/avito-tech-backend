package core

import (
	"avito-tech-backend/internal/core/actions"
	"avito-tech-backend/internal/storage"
	"context"
)

type Repository struct {
	Config  *Config
	Storage *storage.Storage
	Actions *actions.Actions
}

func NewRepository(ctx context.Context, cfg *Config) (*Repository, error) {
	var err error

	r := &Repository{
		Config: cfg,
	}
	r.Storage, err = storage.NewStorage(ctx, cfg.Storage)
	if err != nil {
		return nil, err
	}
	r.Actions = actions.NewActions(r.Storage)
	return r, nil
}
