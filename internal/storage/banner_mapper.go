package storage

import (
	"avito-tech-backend/internal/core/entities"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"strconv"
	"time"
)

type BannerCreateParams struct {
	TagIds    []int64
	FeatureId int64
	Content   string
	IsActive  bool
}

type BannerMapper struct {
	Storage *Storage
}

func (m *BannerMapper) executeQuery(ctx context.Context, query sq.Sqlizer) ([]entities.Banner, error) {
	rows, err := m.Storage.Database.QuerySq(ctx, query)
	if err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]entities.Banner, 0)
	for rows.Next() {
		banner, err := toBanner(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, banner)
	}
	return result, nil
}
func toBanner(rows pgx.Rows) (entities.Banner, error) {
	var banner entities.Banner
	err := rows.Scan(&banner.ID, &banner.TagIds, &banner.FeatureId, &banner.Content, &banner.IsActive, &banner.CreatedAt, &banner.UpdatedAt)
	if err != nil {
		return entities.Banner{}, err
	}
	return banner, nil
}

func (m *BannerMapper) GetAllBannersByTagAndOrFeature(ctx context.Context, tagId int64, featureId int64, limit uint64, offset uint64) ([]entities.Banner, error) {
	return m.executeQuery(ctx, sq.Select("*").From("banners").
		PlaceholderFormat(sq.Dollar).
		Where(sq.And{sq.Eq{"tag_ids": tagId}, sq.Eq{"feature_id": featureId}}).
		Limit(limit).Offset(offset))
}

func (m *BannerMapper) GetBannerByTagAndFeature(ctx context.Context, featureId int64, tagId int64, useLastVersion bool) (*entities.Banner, error) {
	q := sq.Select("*").From("banners").
		PlaceholderFormat(sq.Dollar)
	if useLastVersion {
		q.Where(sq.And{sq.Eq{"feature_id": featureId}, sq.Eq{"tag_ids": tagId}, sq.GtOrEq{"updated_at": time.Now().Add(-time.Minute * 5)}}).
			OrderBy("updated_at DESC").Limit(1)
	} else {
		q.Where(sq.And{sq.Eq{"feature_id": featureId}, sq.Eq{"tag_ids": tagId}})
	}
	q.Where(sq.Eq{"is_active": true})
	result, err := m.executeQuery(ctx, q)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func (m *BannerMapper) InsertBanner(ctx context.Context, params BannerCreateParams) (*entities.Banner, error) {
	result, err := m.executeQuery(ctx, sq.Insert("banners").
		PlaceholderFormat(sq.Dollar).
		Columns("tag_ids", "feature_id", "content", "is_active", "created_at", "updated_at").
		Values(params.TagIds, params.FeatureId, params.Content, params.IsActive, time.Now(), time.Now()).
		Suffix("RETURNING *"))
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func (m *BannerMapper) UpdateBannerById(ctx context.Context, id int64, params entities.RawBanner) (*entities.Banner, error) {
	q := sq.Update("banners").PlaceholderFormat(sq.Dollar)
	if params.TagIds != nil {
		q.Set("tag_ids", params.TagIds)
	}
	if params.FeatureId != "null" {
		featureId, err := strconv.ParseInt(params.FeatureId, 10, 0)
		if err != nil {
			return nil, err
		}
		q.Set("feature_id", featureId)
	}
	if params.Content != "null" {
		q.Set("content", params.Content)
	}
	if params.IsActive != "null" {
		isActive, err := strconv.ParseBool(params.IsActive)
		if err != nil {
			return nil, err
		}
		q.Set("is_active", isActive)
	}
	q.Set("updated_at", time.Now())
	q.Where(sq.Eq{"id": id})
	result, err := m.executeQuery(ctx, q)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func (m *BannerMapper) DeleteBannerById(ctx context.Context, id int64) (*entities.Banner, error) {
	result, err := m.executeQuery(ctx, sq.Delete("banners").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}))
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func (m *BannerMapper) FindBannerById(ctx context.Context, id int64) (*entities.Banner, error) {
	result, err := m.executeQuery(ctx, sq.Select("*").From("banners").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}))
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}

func (m *BannerMapper) FindBannerByTagIdAndFeatureIdAndUseLastVersion(ctx context.Context, tagId int64, featureId int64, useLastVersion bool) (*entities.Banner, error) {
	q := sq.Select("*").From("banners").
		PlaceholderFormat(sq.Dollar)
	if useLastVersion {
		q.Where(sq.And{sq.Eq{"feature_id": featureId}, sq.Eq{"tag_ids": tagId}, sq.GtOrEq{"updated_at": time.Now().Add(-time.Minute * 5)}}).
			OrderBy("updated_at DESC").Limit(1)
	} else {
		q.Where(sq.And{sq.Eq{"feature_id": featureId}, sq.Eq{"tag_ids": tagId}})
	}
	q.Where(sq.Eq{"is_active": true})
	result, err := m.executeQuery(ctx, q)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil
}
