package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/thalysonalexr/learning-go/docker/usecase"
)

func GetMovies(s usecase.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		gender := r.URL.Query().Get("gender")
		movies, err := s.SearchByGender(gender)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error": "error to load movies"}`))
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "error to encode data"}`))
		}
		var toPresenter []MoviePresenter
		for _, movie := range movies {
			formatData, _ := time.Parse(movie.ReleaseNote.String(), "2014-11-12T11:45:26.371Z")
			toPresenter = append(toPresenter, MoviePresenter{
				Title:       movie.Title,
				ReleaseNote: formatData.String(),
			})
		}
		encoded, err := json.Marshal(toPresenter)
		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
	}
}
