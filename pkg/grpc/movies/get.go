package grpc

import (
	"MovieRest/internal/models"
	"context"
	"errors"
	"fmt"
	moviepb "github.com/Olegsuus/MovieProto/gen/models/movie"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (m *MovieRepository) Get(ctx context.Context, id string) (*models.Movie, error) {
	const op = "grpc.Get"

	req := &moviepb.GetRequest{Id: id}

	resp, err := m.client.client.Get(ctx, req)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		log.Printf("%s: %v", op, err)
		return nil, fmt.Errorf("ошибка при получении фильма: %s", err)
	}

	movie := &models.Movie{
		ID:          resp.Movie.Id,
		Title:       resp.Movie.Title,
		Description: resp.Movie.Description,
		Year:        resp.Movie.Year,
		Country:     resp.Movie.Country,
		Genres:      resp.Movie.Genres,
		PosterURL:   resp.Movie.PosterUrl,
		Rating:      resp.Movie.Rating,
	}

	return movie, nil
}
