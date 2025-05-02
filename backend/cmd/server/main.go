package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "your-calendar-app/internal/handler"
    "your-calendar-app/internal/infrastructure"
)

func main() {
    _ = godotenv.Load()

    db, err := infrastructure.NewPostgres(os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("db connect: %v", err)
    }

    r := gin.Default()

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