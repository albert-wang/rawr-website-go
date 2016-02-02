-- +goose Up
CREATE TABLE blog_categories (
	id SERIAL PRIMARY KEY, 
	category VARCHAR(64) NOT NULL
);

CREATE TABLE blog_posts (
	id SERIAL PRIMARY KEY, 
	category INT NOT NULL REFERENCES blog_categories(id),
	title VARCHAR(256) NOT NULL, 
	content TEXT NOT NULL, 
	publish TIMESTAMP
);

CREATE TABLE galleries (
	id SERIAL PRIMARY KEY, 
	name VARCHAR(256) NOT NULL,
	description VARCHAR(2048) NOT NULL,
	displayed BOOLEAN NOT NULL
);

CREATE TABLE gallery_images (
	id SERIAL PRIMARY KEY, 
	gallery INT NOT NULL REFERENCES galleries(id),
	image VARCHAR(2048) NOT NULL,
	name VARCHAR(256) NOT NULL,
	description VARCHAR(2048) NOT NULL,
	publish TIMESTAMP NOT NULL, 
	public BOOLEAN NOT NULL
);

-- +goose Down
DROP TABLE blog_posts;
DROP TABLE blog_categories;
DROP TABLE gallery_images;
DROP TABLE galleries;
