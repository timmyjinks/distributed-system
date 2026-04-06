package task

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	sql *sql.DB
}

func NewService(sql *sql.DB) Service {
	return Service{
		sql: sql,
	}
}

func (s *Service) GetByID(id string) (Task, error) {
	var img Task
	err := s.sql.QueryRow("SELECT * FROM tasks WHERE id = ($1)", id).Scan(&img.ID, &img.Type)
	if err != nil {
		return Task{}, err
	}
	return img, nil
}

func (s *Service) Get() ([]Task, error) {
	var imgs []Task
	rows, err := s.sql.Query("SELECT * FROM tasks")
	if err != nil {
		return []Task{}, err
	}

	for rows.Next() {
		var img Task
		if err := rows.Scan(&img.ID, &img.Type); err != nil {
		}
		imgs = append(imgs, img)
	}

	return imgs, nil
}

func (s *Service) Append(id, t string) (string, error) {
	if _, err := s.sql.Query("INSERT INTO tasks (id, type) VALUES($1, $2) RETURNING id;", id, t); err != nil {
		return "", err
	}

	return id, nil
}

func (s *Service) Job() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	createdAt := time.Now()

	if _, err := s.sql.Query("INSERT INTO jobs (id, type, created_at) VALUES($1, $2, $3);", id, "task", createdAt); err != nil {
		return "", err
	}
	return id.String(), nil
}
