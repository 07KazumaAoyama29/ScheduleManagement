package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "example.com/calendarapp/internal/handler"
    "example.com/calendarapp/internal/infrastructure"
)

func main() {
    _ = godotenv.Load()

    db, err := infrastructure.NewPostgres(os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("db connect: %v", err)
    }

    r := gin.Default()

     // ★ CORS: 開発用に全許可（本番では AllowOrigins を絞ってください）
     r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    if err := r.Run(); err != nil {
        log.Fatalf("server: %v", err)
    }
}