package repositories

import data_models "example/Iris/data-models"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {
}

func NewMovieManager() *MovieManager {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	// 模拟赋值
	movie := &data_models.Movie{
		Name:        "穿越星际",
		Description: "很好看的一部科幻片",
	}
	return movie.Name
}
