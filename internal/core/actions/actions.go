package actions

import (
	"avito-tech-backend/internal/core/entities"
	"avito-tech-backend/internal/storage"
	"context"
)

type Actions struct {
	storage *storage.Storage
}

func NewActions(storage *storage.Storage) *Actions {
	return &Actions{
		storage: storage,
	}
}

func (a *Actions) GetBanner(ctx context.Context, tagId int64, featureId int64, useLastVersion bool) (*entities.Banner, error) {
	return a.storage.Banners.GetBannerByTagAndFeature(ctx, tagId, featureId, useLastVersion)
}

func (a *Actions) GetBanners(ctx context.Context, tagId int64, featureId int64, limit uint64, offset uint64) ([]entities.Banner, error) {
	return a.storage.Banners.GetAllBannersByTagOrFeature(ctx, tagId, featureId, limit, offset)
}

func (a *Actions) UpdateBanner(ctx context.Context, request entities.Banner) (*entities.Banner, error) {
	return a.storage.Banners.UpdateBanner(ctx, request.ID, storage.BannerCreateParams{
		TagIds:    request.TagIds,
		FeatureId: request.FeatureId,
		Content:   request.Content,
		IsActive:  request.IsActive,
	})
}

func (a *Actions) CreateBanner(ctx context.Context, request entities.Banner) (*entities.Banner, error) {
	return a.storage.Banners.InsertBanner(ctx, storage.BannerCreateParams{
		TagIds:    request.TagIds,
		FeatureId: request.FeatureId,
		Content:   request.Content,
		IsActive:  request.IsActive,
	})
}

func (a *Actions) DeleteBanner(ctx context.Context, id int64) (*entities.Banner, error) {
	return a.storage.Banners.DeleteBanner(ctx, id)
}
