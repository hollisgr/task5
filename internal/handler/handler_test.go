package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"task5/internal/model"
	ms "task5/internal/service/mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockService := new(ms.ServiceMock)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	testHandler := New(r, mockService)
	testHandler.Register()

	t.Run("success", func(t *testing.T) {
		expID := 123
		expMovie := model.Movie{}
		reqMovie := CreateMovieRequest{}

		mockService.On("Create", mock.Anything, expMovie).Return(expID, nil).Once()

		body, _ := json.Marshal(reqMovie)
		req, _ := http.NewRequest(http.MethodPost, "/movie", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("bad data", func(t *testing.T) {
		badBody := []byte(`{ "data": "bad data }`)

		req, _ := http.NewRequest(http.MethodPost, "/movie", bytes.NewBuffer(badBody))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestLoad(t *testing.T) {
	mockService := new(ms.ServiceMock)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	testHandler := New(r, mockService)
	testHandler.Register()

	t.Run("success", func(t *testing.T) {
		expID := 123
		expMovie := model.Movie{}

		mockService.On("Load", mock.Anything, expID).Return(expMovie, nil).Once()

		url := fmt.Sprintf("/movie/%d", expID)

		req, _ := http.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("bad id", func(t *testing.T) {
		badUrl := "/movie/badID"

		req, _ := http.NewRequest(http.MethodGet, badUrl, nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("not found", func(t *testing.T) {
		expID := 123
		expMovie := model.Movie{}

		mockService.On("Load", mock.Anything, expID).Return(expMovie, model.ErrNotFound).Once()

		url := fmt.Sprintf("/movie/%d", expID)

		req, _ := http.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		mockService.AssertExpectations(t)
	})
}
