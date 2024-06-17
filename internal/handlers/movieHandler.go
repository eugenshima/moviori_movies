package handlers

import (
	"context"
	"fmt"

	proto "github.com/eugenshima/moviori_movies/proto"
)

type MovieHandler struct {
	hnd MovieService
	proto.UnimplementedMovioriMoviesServer
}

func NewMovieHandler(hnd MovieService) *MovieHandler {
	return &MovieHandler{hnd: hnd}
}

type MovieService interface {
	GetMovieInfo(context.Context, string) ([]byte, error)
}

func (hnd *MovieHandler) FindByName(ctx context.Context, req *proto.FindMovieRequest) (*proto.FindMovieResponse, error) {
	id := req.ID
	movie, err := hnd.hnd.GetMovieInfo(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetMovieInfo: %v", err)
	}
	fmt.Println(string(movie))
	response := &proto.FindMovieResponse{
		MovieInfo: movie,
	}
	return response, nil
}
