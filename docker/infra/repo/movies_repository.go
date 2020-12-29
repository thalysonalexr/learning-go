package repo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/thalysonalexr/learning-go/docker/entity"
	"github.com/thalysonalexr/learning-go/docker/errors"
)

type Repository interface {
	List() ([]entity.Movie, error)
}

type MoviesRepositoryImpl struct{}

func (r *MoviesRepositoryImpl) List() ([]entity.Movie, error) {
	res, err := http.Get("https://raw.githubusercontent.com/meilisearch/MeiliSearch/master/datasets/movies/movies.json")
	if err != nil {
		return []entity.Movie{}, errors.FailedGetResourceMovies
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []entity.Movie{}, errors.FailedToReadResponse
	}
	var movies = []entity.Movie{}
	json.Unmarshal(body, &movies)
	return movies, nil
}
