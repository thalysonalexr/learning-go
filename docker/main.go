package main

import (
	"net/http"

	"github.com/thalysonalexr/learning-go/docker/api"
	repository "github.com/thalysonalexr/learning-go/docker/infra/repo"
	service "github.com/thalysonalexr/learning-go/docker/usecase"
)

func main() {
	repo := repository.MoviesRepositoryImpl{}
	service := service.CreateNewService(&repo)

	http.HandleFunc("/movies", api.GetMovies(service))
	http.ListenAndServe(":8080", nil)
}
