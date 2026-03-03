package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service interface{}

type Handler struct {
	router *gin.Engine
	serv   Service
}

func New(r *gin.Engine, s Service) *Handler {
	return &Handler{
		router: r,
		serv:   s,
	}
}

func (h *Handler) Register() {
	h.router.GET("/hello", h.Hello)
}

func (h *Handler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
