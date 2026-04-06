-- +goose Up
CREATE TABLE images(
  id uuid PRIMARY KEY,
  name text,
  content text
);

-- +goose Down
DROP TABLE images
