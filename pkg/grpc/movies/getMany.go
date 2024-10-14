package grpc

import (
	"context"
	"fmt"
	moviepb "github.com/Olegsuus/MovieProto/gen/models/movie"
	"github.com/Olegsuus/MovieRest/internal/models"
	"log"
)

func (m *MovieRepository) GetMany(ctx context.Context) ([]*models.Movie, error) {
	const op = "grpc.GetMany"

	req := &moviepb.GetManyRequest{}

	resp, err := m.client.client.GetMany(ctx, req)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return nil, fmt.Errorf("ошибка при получении списка фильмов: %w", err)
	}

	var movies []*models.Movie
	for _, grpcMovie := range resp.Movies {
		movie := &models.Movie{
			ID:          grpcMovie.Id,
			Title:       grpcMovie.Title,
			Description: grpcMovie.Description,
			Year:        grpcMovie.Year,
			Country:     grpcMovie.Country,
			Genres:      grpcMovie.Genres,
			PosterURL:   grpcMovie.PosterUrl,
			Rating:      grpcMovie.Rating,
		}
		movies = append(movies, movie)
	}

	return movies, nil
}
