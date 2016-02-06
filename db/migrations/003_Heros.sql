-- +goose Up
ALTER TABLE blog_posts ADD COLUMN hero TEXT NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE blog_posts DROP COLUMN hero;