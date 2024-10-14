package grpc

import (
	"context"
	"fmt"
	moviepb "github.com/Olegsuus/MovieProto/gen/models/movie"
	"log"
)

func (m *MovieRepository) Remove(ctx context.Context, id string) error {
	const op = "grpc.remove"

	req := &moviepb.RemoveRequest{
		Id: id,
	}

	resp, err := m.client.client.Remove(ctx, req)
	if err != nil {
		log.Printf("%s:%v", op, err)
		return fmt.Errorf("ошибка при удалении фильма из базы: %s", err)
	}

	if !resp.Status {
		return fmt.Errorf("не удалось удалить фильм с ID %s", id)
	}

	return nil

}
