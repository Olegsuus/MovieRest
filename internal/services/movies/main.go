package services

import (
	"MovieRest/internal/models"
	"context"
	"log/slog"
)

type MoviesService struct {
	msP MoviesServiceProvider
	l   *slog.Logger
}

type MoviesServiceProvider interface {
	Add(ctx context.Context, movie *models.Movie) (string, error)
	Get(ctx context.Context, id string) (*models.Movie, error)
	GetMany(ctx context.Context) ([]*models.Movie, error)
	Update(ctx context.Context, id string, movie *models.Movie) error
	Remove(ctx context.Context, id string) error
}

func NewMoviesService(msP MoviesServiceProvider, logger *slog.Logger) *MoviesService {
	return &MoviesService{
		msP: msP,
		l:   logger,
	}
}
