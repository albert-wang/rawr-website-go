-- +goose Up
ALTER TABLE galleries ADD COLUMN hero TEXT NOT NULL DEFAULT '';
ALTER TABLE galleries ADD COLUMN s3prefix TEXT NOT NULL DEFAULT '';
DROP TABLE gallery_images;

-- +goose Down
ALTER TABLE galleries DROP COLUMN hero;
CREATE TABLE gallery_images (
	id SERIAL PRIMARY KEY, 
	gallery INT NOT NULL REFERENCES galleries(id),
	image VARCHAR(2048) NOT NULL,
	name VARCHAR(256) NOT NULL,
	description VARCHAR(2048) NOT NULL,
	publish TIMESTAMP NOT NULL, 
	public BOOLEAN NOT NULL
);