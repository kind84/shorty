package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/kind84/shorty/api"
	"github.com/kind84/shorty/db"
	"github.com/kind84/shorty/usecase"
)

func main() {
	db, err := db.NewRedisDB("redis")
	if err != nil {
		log.Fatalf("error initializing database: %s", err)
	}
	app := usecase.NewApp(db)
	service := api.NewService(app)

	r := gin.Default()
	r.GET("/:hash", service.Bounce)
	r.POST("/api/cut", service.Cut)
	r.POST("/api/destroy", service.Burn)
	r.POST("/api/inflate", service.Inflate)

	r.Run()
}
