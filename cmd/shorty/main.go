package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kind84/shorty/api"
	"github.com/kind84/shorty/usecase"
)

func main() {
	app := usecase.NewApp()
	service := api.NewService(app)

	r := gin.Default()
	r.GET("/:hash", service.Bounce)
	r.POST("/api/cut", service.Cut)
	r.POST("/api/destroy", service.Burn)

	r.Run()
}
