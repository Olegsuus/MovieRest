package handlers

import (
	"MovieRest/internal/models"
	"context"
)

type MovieHandlers struct {
	mhP MovieHandlersProvider
}

type MovieHandlersProvider interface {
	Add(ctx context.Context, movie *AddMovieDTO) (string, error)
	Get(ctx context.Context, id string) (*models.Movie, error)
	GetMany(ctx context.Context) ([]*models.Movie, error)
	Remove(ctx context.Context, id string) error
	Update(ctx context.Context, id string, movie *models.Movie) error
}

func NewMovieHandlers(mhP MovieHandlersProvider) *MovieHandlers {
	return &MovieHandlers{
		mhP: mhP,
	}
}
