CREATE TABLE IF NOT EXISTS banners
(
    id          serial      PRIMARY KEY,
    tag_ids     integer[]        NOT NULL,
    feature_id  int         NOT NULL,
    content     string      NOT NULL,
    is_active   boolean     NOT NULL,
    created_at  timestamp   NOT NULL,
    updated_at  timestamp   NOT NULL
);

CREATE INDEX idx_tags_ids ON banners USING gin(tag_ids)