package mem

import (
	"errors"
	"fmt"
	"grpc-microservice-example/models"
)

var MovieMem map[string]models.Movie

func GetMovie(id string) (models.Movie, error) {
	if val, ok := MovieMem[id]; ok {
		return val, nil
	}
	return models.Movie{}, errors.New("movie not found in memory")
}

func UpdateMovie(m models.Movie) {
	fmt.Printf("%v", m)
	MovieMem[m.ID] = m
}

func DeleteMovie(id string) {
	delete(MovieMem, id)
}

func AddMoview(m models.Movie) {
	MovieMem[m.ID] = m
}
