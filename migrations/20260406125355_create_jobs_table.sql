-- +goose Up
CREATE TABLE jobs(
  id uuid PRIMARY KEY,
  type text,
  created_at TIMESTAMPTZ 
);

-- +goose Down
DROP TABLE jobs 
