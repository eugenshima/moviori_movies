package repository

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type KinopoiskApiRepo struct {
}

func NewKinopoiskApiRepo() *KinopoiskApiRepo {
	return &KinopoiskApiRepo{}
}

func (kino *KinopoiskApiRepo) FindByID(ctx context.Context, id string) ([]byte, error) {
	url := "https://api.kinopoisk.dev/v1.4/movie/" + id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %v", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", "YDQGDCE-8AFMY5N-GEDFY02-14XCVG7")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.DefaultClient.Do: %v", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %v", err)
	}

	// movie := &model.Movie{}
	// err := json.Unmarshal(body, &movie)
	// if err != nil {
	// 	return nil, fmt.Errorf(" Unmarshal: %w", err)
	// }
	return body, nil
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
