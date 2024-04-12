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

func (a *Actions) GetUserBanner(ctx context.Context, tagId int64, featureId int64, useLastVersion bool) (*entities.Banner, error) {
	return a.storage.Banners.GetBannerByTagAndFeature(ctx, tagId, featureId, useLastVersion)
}

func (a *Actions) GetBanners(ctx context.Context, tagId int64, featureId int64, limit uint64, offset uint64) ([]entities.Banner, error) {
	return a.storage.Banners.GetAllBannersByTagAndOrFeature(ctx, tagId, featureId, limit, offset)
}

func (a *Actions) UpdateBanner(ctx context.Context, request entities.Banner) (*entities.Banner, error) {
	banner, err := a.storage.Banners.FindBannerById(ctx, request.ID)
	if err != nil {
		return nil, err
	}
	if banner == nil {
		return nil, nil
	}
	return a.storage.Banners.UpdateBannerById(ctx, request.ID, storage.BannerCreateParams{
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
	banner, err := a.storage.Banners.FindBannerById(ctx, id)
	if err != nil {
		return nil, err
	} else if banner == nil {
		return nil, nil
	}
	return a.storage.Banners.DeleteBannerById(ctx, id)
}
