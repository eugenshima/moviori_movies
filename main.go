package main

import (
	"context"
	"fmt"

	repo "github.com/eugenshima/moviori_movies/internal/repository"
)

func main() {

	rps := repo.NewKinopoiskApiRepo()
	value, err := rps.FindByID(context.Background(), "100000")
	if err != nil {
		fmt.Errorf("FindByID: %v", err)
	}

	fmt.Println("Название фильма -", value.AlternativeName, "\nГод выхода -", value.Year, "\nПродолжительность -", value.MovieLength, "минут")

}
