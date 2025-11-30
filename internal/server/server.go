package server

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

type Server struct {
    router *gin.Engine
    port   string
}

// NewServer initializes the Gin router and returns a Server instance
func NewServer(port string) *Server {
    r := gin.Default()

    // Example health check route
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "ok",
        })
    })

    return &Server{
        router: r,
        port:   port,
    }
}

// Start runs the server
func (s *Server) Start() error {
    fmt.Println("Server running on port", s.port)
    return s.router.Run(":" + s.port)
}
