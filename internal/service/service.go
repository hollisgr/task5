package service

import (
	"context"
	"sync"
	"task5/internal/model"
)

type Service struct {
	db []model.Movie
	mu sync.Mutex
}

func New() *Service {
	return &Service{
		db: make([]model.Movie, 0),
	}
}

func (s *Service) Create(ctx context.Context, data model.Movie) (int, error) {
	if data.Rating < 0 {
		return 0, model.ErrInvalidRating
	}

	if s.db == nil {
		return 0, model.ErrDBInternal
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	id := len(s.db) + 1
	data.ID = id
	s.db = append(s.db, data)
	return id, nil
}

func (s *Service) Load(ctx context.Context, id int) (model.Movie, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if id > len(s.db) {
		return model.Movie{}, model.ErrNotFound
	}

	return s.db[id-1], nil
}
