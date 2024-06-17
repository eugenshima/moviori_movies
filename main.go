package main

import (
	"fmt"
	"log"
	"net"

	"github.com/eugenshima/moviori_movies/internal/handlers"
	repo "github.com/eugenshima/moviori_movies/internal/repository"
	"github.com/eugenshima/moviori_movies/internal/service"
	pb "github.com/eugenshima/moviori_movies/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	rps := repo.NewKinopoiskApiRepo()
	srv := service.NewMovieService(rps)
	hnd := handlers.NewMovieHandler(srv)

	s := grpc.NewServer()
	pb.RegisterMovioriMoviesServer(s, hnd)
	err = s.Serve(lis)
	if err != nil {
		fmt.Errorf("cannot start server: %v", err)
	}

}
