package main

import (
	"context"
	"log"
	"os"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"github.com/kind84/shorty/api"
	"github.com/kind84/shorty/db"
	"github.com/kind84/shorty/usecase"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic trapped in main goroutine : %+v", err)
			log.Printf("stacktrace from panic: %s", string(debug.Stack()))
			os.Exit(1)
		}
	}()

	ctx := context.Background()

	db, err := db.NewRedisDB(ctx, "redis:6379")
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
	r.POST("/api/count", service.Count)

	r.Run()
}
