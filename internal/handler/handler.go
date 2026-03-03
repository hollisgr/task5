package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"task5/internal/model"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Create(ctx context.Context, data model.Movie) (int, error)
	Load(ctx context.Context, id int) (model.Movie, error)
}

type Handler struct {
	router  *gin.Engine
	service Service
}

func New(r *gin.Engine, s Service) *Handler {
	return &Handler{
		router:  r,
		service: s,
	}
}

func (h *Handler) Register() {
	h.router.POST("/movie", h.Create)
	h.router.GET("/movie/:id", h.Load)
}

func (h *Handler) Create(c *gin.Context) {
	req := CreateMovieRequest{}
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid body",
		})
		return
	}
	data := req.ToModel()
	id, err := h.service.Create(c.Request.Context(), data)
	if err != nil {
		if err == model.ErrDBInternal {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "internal err",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"id":      id,
	})
}

func (h *Handler) Load(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid id",
		})
		return
	}

	res, err := h.service.Load(c.Request.Context(), id)
	if err != nil {
		if err == model.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "internal err",
			})
			return
		}
		// другие ошибки не сервис не возращает
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}
