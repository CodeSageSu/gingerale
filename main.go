package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Health check route
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    // Simple welcome route
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Welcome to gingerale with Gin!")
    })

    // Starts the server on 0.0.0.0:8080
    r.Run()
}

