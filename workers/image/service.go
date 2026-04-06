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
	"github.com/timmyjinks/distributed-system/utils"
)

type Service struct {
	sql   *sql.DB
	queue *queue.KafkaService
}

func NewService(ctx context.Context, sql *sql.DB, q *queue.KafkaService) Service {
	service := Service{
		sql:   sql,
		queue: q,
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("[INFO] image worker shutting down")
				return
			default:
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

				go func() {
					utils.SimulateLargeTask(time.Second * 10)

					id, err := service.Job()
					if err != nil {
						log.Println("[WARN] job failed")
						return
					}

					if _, err := service.Append(id, payload.Name, payload.Content); err != nil {
						log.Println("[WARN] Append failed")
						return
					}
				}()
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
	_, err := s.sql.Exec("INSERT INTO images (id, name, content) VALUES($1, $2, $3);", id, name, content)
	if err != nil {
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

	if _, err := s.sql.Exec("INSERT INTO jobs (id, type, created_at) VALUES($1, $2, $3);", id, "image", createdAt); err != nil {
		return "", err
	}

	return id.String(), nil
}
