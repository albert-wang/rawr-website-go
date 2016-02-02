-- +goose Up
INSERT INTO blog_categories (category) VALUES ('uncategorized');
INSERT INTO blog_categories (category) VALUES ('featured');
INSERT INTO blog_categories (category) VALUES ('programming');
INSERT INTO blog_categories (category) VALUES ('araboth');
INSERT INTO blog_categories (category) VALUES ('random');

INSERT INTO galleries (name, description, displayed) VALUES ('Araboth', '', TRUE);
INSERT INTO galleries (name, description, displayed) VALUES ('random', 'random uploads', FALSE);


-- +goose Down
DELETE FROM blog_categories;
DELETE FROM galleries;