package services

import (
	"context"
	"fmt"
	"github.com/Olegsuus/MoviesRest/internal/models"
	"log/slog"
)

func (s *MoviesService) GetMany(ctx context.Context) ([]*models.Movie, error) {
	const op = "services.GetMany"

	s.l.With(slog.String("op", op))

	movies, err := s.msP.GetMany(ctx)
	if err != nil {
		s.l.Error("ошибка при получении списка фильмов", slog.Any("error", err))
		return nil, fmt.Errorf("ошибка при получении списка фильмов: %w", err)
	}

	s.l.Info("Успешное получение списка фильмов")

	return movies, nil
}
