package grpc

import (
	"context"
	"fmt"
	moviepb "github.com/Olegsuus/MovieProto/gen/models/movie"
	"github.com/Olegsuus/MoviesRest/internal/models"
	"log"
)

func (m *MovieRepository) Add(ctx context.Context, movie *models.Movie) (string, error) {
	const op = "grpc.Add"

	pbMovie := &moviepb.Movie{
		Title:       movie.Title,
		Description: movie.Description,
		Year:        movie.Year,
		Country:     movie.Country,
		Genres:      movie.Genres,
		PosterUrl:   movie.PosterURL,
		Rating:      movie.Rating,
	}

	req := &moviepb.AddRequest{Movie: pbMovie}

	resp, err := m.client.client.Add(ctx, req)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return "", fmt.Errorf("ошибка при добавлении фильма: %s", err)
	}

	return resp.Id, nil
}
