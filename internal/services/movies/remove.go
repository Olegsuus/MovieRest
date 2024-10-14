package services

import (
	"context"
	"fmt"
	"log/slog"
)

func (s *MoviesService) Remove(ctx context.Context, id string) error {
	const op = "services.Remove"

	s.l.With(slog.String("op", op))

	err := s.msP.Remove(ctx, id)
	if err != nil {
		s.l.Error("ошибка при удалении фильма", slog.Any("error", err))
		return fmt.Errorf("ошибка при удалении фильма: %w", err)
	}

	s.l.Info("Успешное удаление фильма")

	return nil
}
