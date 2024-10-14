package grpc

import (
	moviepb "github.com/Olegsuus/MovieProto/gen/models/movie"
	"google.golang.org/grpc"
)

type MovieClient struct {
	conn   *grpc.ClientConn
	client moviepb.MovieServiceClient
}

func NewMovieClient(address string) (*MovieClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := moviepb.NewMovieServiceClient(conn)

	return &MovieClient{
		conn:   conn,
		client: client,
	}, nil
}

func (c *MovieClient) Close() error {
	return c.conn.Close()
}
