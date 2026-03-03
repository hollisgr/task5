package service

import (
	"context"
	"task5/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	service := New()

	t.Run("success", func(t *testing.T) {
		movie := model.Movie{}
		expID := 1
		id, err := service.Create(context.Background(), movie)
		assert.NoError(t, err)
		assert.Equal(t, expID, id)
	})

	t.Run("invalid rating", func(t *testing.T) {
		movie := model.Movie{
			Rating: -123,
		}
		expID := 0
		expErr := model.ErrInvalidRating
		id, err := service.Create(context.Background(), movie)
		assert.Error(t, err)
		assert.Equal(t, expID, id)
		assert.Equal(t, expErr, err)
	})
}

func TestLoad(t *testing.T) {
	service := New()
	testMovie := model.Movie{
		Title:       "test_title",
		Rating:      123,
		Genre:       "test_genre",
		Description: "test_description",
	}
	id, _ := service.Create(context.Background(), testMovie) // добавляем в базу тест муви
	testMovie.ID = id

	t.Run("success", func(t *testing.T) {
		expMovie := testMovie
		res, err := service.Load(context.Background(), id)
		assert.NoError(t, err)
		assert.Equal(t, expMovie, res)
	})

	t.Run("fail", func(t *testing.T) {
		expMovie := model.Movie{}
		expErr := model.ErrNotFound
		res, err := service.Load(context.Background(), 123)
		assert.Error(t, err)
		assert.Equal(t, expMovie, res)
		assert.Equal(t, expErr, err)
	})
}
