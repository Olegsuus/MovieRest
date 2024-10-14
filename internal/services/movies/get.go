package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/Olegsuus/MovieRest/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

func (s *MoviesService) Get(ctx context.Context, id string) (*models.Movie, error) {
	const op = "services.Get"

	s.l.With(slog.String("op", op))

	movie, err := s.Get(ctx, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}

		s.l.Error("ошибка при получении фильма", op, err)
		return nil, fmt.Errorf("ошибка при получении фильма: %s", err)
	}

	s.l.Info("Успешное получение фильма")

	return movie, nil
}
