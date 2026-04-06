package image

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/timmyjinks/distributed-system/queue"
	// "github.com/timmyjinks/distributed-system/utils"
)

type Service struct {
	sql   *sql.DB
	queue *queue.KafkaService
}

func NewService(sql *sql.DB, q *queue.KafkaService) Service {
	service := Service{
		sql:   sql,
		queue: q,
	}

	go func() {
		for {
			msg, err := q.Consumer.Read(context.Background())
			if err != nil {
				log.Println("[WARN] Consumer read Bad data")
				continue
			}
			fmt.Println("consumed", msg)
			var payload Image
			if err := json.Unmarshal(msg.Payload, &payload); err != nil {
				log.Println("[WARN] Bad data")
				continue
			}

			// utils.SimulateLargeTask(time.Second * 10)

			id, err := service.Job()
			if err != nil {
				log.Println("[WARN] job failed")
				continue
			}

			if _, err := service.Append(id, payload.Name, payload.Content); err != nil {
				log.Println("[WARN] Append failed")
				continue
			}
		}
	}()

	return service
}

func (s *Service) GetByID(id string) (Image, error) {
	var img Image
	err := s.sql.QueryRow("SELECT * FROM images WHERE id = ($1);", id).Scan(&img.ID, &img.Name, &img.Content)
	if err != nil {
		return Image{}, err
	}
	return img, nil
}

func (s *Service) Get() ([]Image, error) {
	var imgs []Image
	rows, err := s.sql.Query("SELECT * FROM images")
	if err != nil {
		return []Image{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var img Image
		if err := rows.Scan(&img.ID, &img.Name, &img.Content); err != nil {
		}
		imgs = append(imgs, img)
	}

	return imgs, nil
}

func (s *Service) Append(id, name, content string) (string, error) {
	res, err := s.sql.Query("INSERT INTO images (id, name, content) VALUES($1, $2, $3);", id, name, content)
	if err != nil {
		return "", err
	}
	defer res.Close()

	return id, nil
}

func (s *Service) Job() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	createdAt := time.Now()

	res, err := s.sql.Query("INSERT INTO jobs (id, type, created_at) VALUES($1, $2, $3);", id, "image", createdAt)
	if err != nil {
		return "", err
	}
	defer res.Close()

	return id.String(), nil
}
