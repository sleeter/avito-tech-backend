package handlers

import (
	"avito-tech-backend/internal/core"
	"avito-tech-backend/internal/core/entities"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func GetUserBanner(ctx *gin.Context, r *core.Repository) error {
	var queryParams = struct {
		TagId          int64 `form:"tag_id" required:"true"`
		FeatureId      int64 `form:"feature_id" required:"true"`
		UseLastVersion bool  `form:"use_last_version"`
	}{
		UseLastVersion: false,
	}
	if err := ctx.BindQuery(&queryParams); err != nil {
		slog.Debug("Error with getting user banner: %s", err)
		//zap.L().Debug("Get user banner", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil
	}
	banner, err := r.Actions.GetUserBanner(ctx, queryParams.TagId, queryParams.FeatureId, queryParams.UseLastVersion)
	if err != nil {
		return err
	}
	if banner == nil {
		slog.Debug("Error with getting user banner: %s", err)
		ctx.Status(http.StatusNotFound)
		return nil
	}
	ctx.JSON(http.StatusOK, gin.H{
		"content": banner.Content,
	})
	return nil
}

func CreateBanner(ctx *gin.Context, r *core.Repository) error {
	var Banner struct {
		TagIds    []int64 `json:"tag_ids" required:"true"`
		FeatureId int64   `json:"feature_id" required:"true"`
		Content   string  `json:"content" required:"true"`
		IsActive  bool    `json:"is_active" required:"true"`
	}
	if err := ctx.BindJSON(&Banner); err != nil {
		slog.Debug("Error with creating banner: %s", err)
		//zap.L().Debug("Create courier", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil
	}

	banner, err := r.Actions.CreateBanner(ctx, entities.Banner{
		TagIds:    Banner.TagIds,
		FeatureId: Banner.FeatureId,
		Content:   Banner.Content,
		IsActive:  Banner.IsActive,
	})
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"banner_id": banner.ID,
	})
	return nil
}

func UpdateBanner(ctx *gin.Context, r *core.Repository) error {
	bannerId, err := strconv.ParseInt(ctx.Param("banner_id"), 10, 0)
	if err != nil {
		slog.Debug("Error with updating banner: %s", err)
		//zap.L().Debug("Update banner", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil
	}
	var Banner struct {
		TagIds    string `json:"tag_ids"`
		FeatureId string `json:"feature_id"`
		Content   string `json:"content"`
		IsActive  string `json:"is_active"`
	}
	if err := ctx.BindJSON(&Banner); err != nil {
		slog.Debug("Error with updating banner: %s", err)
		//zap.L().Debug("Create courier", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil
	}
	if Banner.TagIds == "null" && Banner.FeatureId == "null" && Banner.Content == "null" && Banner.IsActive == "null" {
		slog.Debug("Error with updating banner: you must pass at least one parameter")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error with updating banner: you must pass at least one parameter",
		})
		return nil
	}
	var TagIdsStruct struct {
		TagIds []int64 `json:"tag_ids"`
	}
	if Banner.TagIds != "null" {
		if err := ctx.BindJSON(&TagIdsStruct); err != nil {

		}
	} else {
		TagIdsStruct.TagIds = nil
	}
	banner, err := r.Actions.UpdateBanner(ctx, entities.RawBanner{
		ID:        bannerId,
		TagIds:    TagIdsStruct.TagIds,
		FeatureId: Banner.FeatureId,
		Content:   Banner.Content,
		IsActive:  Banner.IsActive,
	})
	if err != nil {
		return err
	}
	if banner == nil {
		ctx.Status(http.StatusNotFound)
		return nil
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return nil
}

func DeleteBanner(ctx *gin.Context, r *core.Repository) error {
	bannerId, err := strconv.ParseInt(ctx.Param("banner_id"), 10, 0)
	if err != nil {
		slog.Debug("Error with deleting banner: %s", err)
		//zap.L().Debug("Delete banner", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil
	}
	banner, err := r.Actions.DeleteBanner(ctx, bannerId)
	if err != nil {
		return err
	}
	if banner == nil {
		ctx.Status(http.StatusNotFound)
		return nil
	}
	ctx.Status(http.StatusNoContent)
	return nil
}

func GetBanners(ctx *gin.Context, r *core.Repository) error {
	var queryParams = struct {
		TagId     int64  `form:"tag_id"`
		FeatureId int64  `form:"feature_id"`
		Limit     uint64 `form:"limit"`
		Offset    uint64 `form:"offset"`
	}{
		Limit:  1,
		Offset: 0,
	}
	if err := ctx.BindQuery(&queryParams); err != nil {
		slog.Debug("Error with getting banners: %s", err)
		//zap.L().Debug("Get banner", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil
	}
	banners, err := r.Actions.GetBanners(ctx, queryParams.TagId, queryParams.FeatureId, queryParams.Limit, queryParams.Offset)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, banners)
	return nil
}
