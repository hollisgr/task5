package handler

import "task5/internal/model"

type CreateMovieRequest struct {
	Title       string  `json:"title"`
	Rating      float64 `json:"rating"`
	Genre       string  `json:"genre"`
	Description string  `json:"description"`
}

func (r *CreateMovieRequest) ToModel() model.Movie {
	return model.Movie{
		Title:       r.Title,
		Rating:      r.Rating,
		Genre:       r.Genre,
		Description: r.Description,
	}
}
