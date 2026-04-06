-- +goose Up
CREATE TABLE tasks(
  id uuid PRIMARY KEY,
  type text
);

-- +goose Down
DROP TABLE tasks 
