package grpc

import (
	"context"
	"fmt"
	moviepb "github.com/Olegsuus/MovieProto/gen/models/movie"
	"github.com/Olegsuus/MoviesRest/internal/models"
	"log"
)

func (m *MovieRepository) Update(ctx context.Context, id string, movie *models.Movie) error {
	const op = "grpc.Update"

	pbMovie := &moviepb.Movie{
		Id:          id,
		Title:       movie.Title,
		Description: movie.Description,
		Year:        movie.Year,
		Country:     movie.Country,
		Genres:      movie.Genres,
		PosterUrl:   movie.PosterURL,
		Rating:      movie.Rating,
	}

	req := &moviepb.UpdateRequest{
		Movie: pbMovie,
	}

	resp, err := m.client.client.Update(ctx, req)
	if err != nil {
		log.Printf("%s: %v", op, err)
		return fmt.Errorf("ошибка при обновлении фильма: %w", err)
	}

	if !resp.Status {
		return fmt.Errorf("не удалось обновить фильм с ID %s", id)
	}

	return nil
}
