package grpc

import (
	"context"
	"github.com/Olegsuus/MoviesRest/internal/models"
)

type MovieRepository struct {
	client *MovieClient
}

type MovieProvider interface {
	Add(ctx context.Context, movie *models.Movie) (string, error)
	Get(ctx context.Context, id string) (*models.Movie, error)
	GetMany(ctx context.Context) ([]*models.Movie, error)
	Update(ctx context.Context, id string, movie *models.Movie) error
	Remove(ctx context.Context, id string) error
}

func NewMovieRepository(client *MovieClient) *MovieRepository {
	return &MovieRepository{
		client: client,
	}
}
