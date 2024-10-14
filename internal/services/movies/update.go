package services

import (
	"MovieRest/internal/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *MoviesService) Update(ctx context.Context, id string, movie *models.Movie) error {
	const op = "services.Update"

	s.l.With(slog.String("op", op))

	err := s.msP.Update(ctx, id, movie)
	if err != nil {
		s.l.Error("ошибка при обновлении фильма", slog.Any("error", err))
		return fmt.Errorf("ошибка при обновлении фильма: %w", err)
	}

	s.l.Info("Успешное обновление фильма")

	return nil
}
