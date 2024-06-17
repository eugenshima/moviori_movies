package service

import (
	"context"
	"fmt"
)

type MovieService struct {
	srv MovieRepo
}

func NewMovieService(srv MovieRepo) *MovieService {
	return &MovieService{srv: srv}
}

type MovieRepo interface {
	FindByID(context.Context, string) ([]byte, error)
}

func (s *MovieService) GetMovieInfo(ctx context.Context, id string) ([]byte, error) {
	movie, err := s.srv.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("FindByID: %v", err)
	}
	return movie, nil
}
