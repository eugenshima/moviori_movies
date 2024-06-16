package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/eugenshima/moviori_movies/internal/model"
)

type KinopoiskApiRepo struct {
}

func NewKinopoiskApiRepo() *KinopoiskApiRepo {
	return &KinopoiskApiRepo{}
}

func (kino *KinopoiskApiRepo) FindByID(ctx context.Context, id string) (*model.Movie, error) {
	url := "https://api.kinopoisk.dev/v1.4/movie/" + id

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", "YDQGDCE-8AFMY5N-GEDFY02-14XCVG7")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	movie := &model.Movie{}
	err := json.Unmarshal(body, &movie)
	if err != nil {
		return nil, fmt.Errorf(" Unmarshal: %w", err)
	}
	return movie, nil
}

// func (kino *KinopoiskApiRepo) FindByName(ctx context.Context, id string) (*model.movieByName, error) {
// 	//encodeName := encodeURL("Трансформеры")
// 	url := "https://api.kinopoisk.dev/v1.4/movie/search?page=1&limit=1&query=the%20fast%20and%20the%20furious"

// 	req, _ := http.NewRequest("GET", url, nil)

// 	req.Header.Add("accept", "application/json")
// 	req.Header.Add("X-API-KEY", "YDQGDCE-8AFMY5N-GEDFY02-14XCVG7")

// 	res, _ := http.DefaultClient.Do(req)

// 	defer res.Body.Close()
// 	body, _ := io.ReadAll(res.Body)
// 	fmt.Println(string(body))
// 	movie := &model.Response{}
// 	err := json.Unmarshal(body, &movie)
// 	if err != nil {
// 		return nil, fmt.Errorf(" Unmarshal: %w", err)
// 	}
// 	movieByName := &model.MovieByName{}
// 	err = json.Unmarshal([]byte(movie.Docs), &movieByName)
// 	if err != nil {
// 		return nil, fmt.Errorf(" Unmarshal: %w", err)
// 	}
// 	return movieByName, nil
// }
