package services

import (
	"context"
	"fmt"
	handlers "github.com/Olegsuus/MoviesRest/internal/handlers/movies"
	"log/slog"
)

func (s *MoviesService) Add(ctx context.Context, movieDTO *handlers.AddMovieDTO) (string, error) {
	const op = "services.Add"

	s.l.With(slog.String("op", op))

	movie, err := s.SearchInfoForMovie(movieDTO.Title)
	if err != nil {
		s.l.Error("ошибка при поиске фильма во внутреннем api", op, err)
		return "", fmt.Errorf("ошибка при поиске фильма во втрутненнем api")
	}

	id, err := s.msP.Add(ctx, movie)
	if err != nil {
		s.l.Error("ошибка при добавлении нового фильма", op, err)
		return "", fmt.Errorf("ошибка при добавлении нового фильма")
	}

	s.l.Info("Успешное добавление фильма")

	return id, nil
}
