package entities

import "time"

type Banner struct {
	ID        int64      `json:"banner_id"`
	TagIds    []int64    `json:"tag_ids"`
	FeatureId int64      `json:"feature_id"`
	Content   string     `json:"content"`
	IsActive  bool       `json:"is_active"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
type RawBanner struct {
	ID        int64   `json:"banner_id"`
	TagIds    []int64 `json:"tag_ids"`
	FeatureId string  `json:"feature_id"`
	Content   string  `json:"content"`
	IsActive  string  `json:"is_active"`
}
