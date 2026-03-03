package main

import (
	"net/http"
	"task5/internal/app"
	"task5/internal/handler"
	"task5/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	s := service.New()

	r := gin.Default()

	srv := &http.Server{
		Addr:    "127.0.0.1:8008",
		Handler: r,
	}

	h := handler.New(r, s)
	h.Register()

	app.StartServer(srv)

	app.HandleQuit(srv)
}
