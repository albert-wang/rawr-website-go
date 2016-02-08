-- +goose Up
ALTER TABLE blog_categories ADD COLUMN hidden BOOLEAN NOT NULL DEFAULT false;

-- +goose Down
ALTER TABLE blog_categories DROP COLUMN hidden;