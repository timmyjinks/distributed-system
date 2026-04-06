-- +goose Up
CREATE TABLE reports(
  id uuid PRIMARY KEY,
  title text,
  body text
);

-- +goose Down
DROP TABLE reports 
