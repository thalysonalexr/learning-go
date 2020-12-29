package usecase

import (
	"strings"

	"github.com/thalysonalexr/learning-go/docker/entity"
	"github.com/thalysonalexr/learning-go/docker/infra/repo"
)

// Service interface of usecase
type Service interface {
	SearchByGender(k string) ([]entity.Movie, error)
}

// ServiceImpl is implementation of Service interface
type ServiceImpl struct {
	repo repo.Repository
}

// CreateNewService is a factory to create new Service
func CreateNewService(r repo.Repository) *ServiceImpl {
	return &ServiceImpl{
		repo: r,
	}
}

// SearchByGender is a handler to get list movies
func (service *ServiceImpl) SearchByGender(k string) ([]entity.Movie, error) {
	movies, err := service.repo.List()
	if err != nil {
		return []entity.Movie{}, err
	}
	filtered := []entity.Movie{}
	for _, movie := range movies {
		for _, gender := range movie.Genres {
			if strings.Contains(strings.ToLower(gender), strings.ToLower(k)) {
				filtered = append(filtered, movie)
			}
		}
	}
	return filtered, nil
}
