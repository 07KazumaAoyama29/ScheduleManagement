package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"example.com/calendarapp/internal/handler"
	"example.com/calendarapp/internal/infrastructure"
)

func main() {
	_ = godotenv.Load()

	// DB 接続
	db, err := infrastructure.NewPostgres(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}

	// Gin + CORS
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	// ルーティング
	h := handler.NewEventHandler(db)
	api := r.Group("/api")
	{
		api.GET("/events", h.List)
		api.POST("/events", h.Create)
	}

	if err := r.Run(); err != nil {
		log.Fatalf("server: %v", err)
	}
}