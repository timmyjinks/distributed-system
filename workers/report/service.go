package report

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

func (s *Service) GetByID(id string) (Report, error) {
	var img Report
	err := s.sql.QueryRow("SELECT * FROM reports WHERE id = ($1)", id).Scan(&img.ID, &img.Title, &img.Body)
	if err != nil {
		return Report{}, err
	}
	return img, nil
}

func (s *Service) Get() ([]Report, error) {
	var imgs []Report
	rows, err := s.sql.Query("SELECT * FROM reports")
	if err != nil {
		return []Report{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var img Report
		if err := rows.Scan(&img.ID, &img.Title, &img.Body); err != nil {
		}
		imgs = append(imgs, img)
	}

	return imgs, nil
}

func (s *Service) Append(id, title, body string) (string, error) {
	if _, err := s.sql.Exec("INSERT INTO reports (id, title, body) VALUES($1, $2, $3) RETURNING id;", id, title, body); err != nil {
		return "", err
	}

	return id, nil
}

func (s *Service) Job() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	createdAt := time.Now()

	if _, err := s.sql.Exec("INSERT INTO jobs (id, type, created_at) VALUES($1, $2, $3);", id, "report", createdAt); err != nil {
		return "", err
	}
	return id.String(), nil
}
