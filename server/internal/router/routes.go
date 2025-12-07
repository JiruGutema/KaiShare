// Package routes contains all route definitions for the API server
package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jirugutema/kaishare/internal/handler"
	"github.com/jirugutema/kaishare/internal/middleware"
)

func Routes() *gin.Engine {
	router := gin.Default()
	// router.Use(pkg.Logger)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "https://kaishare.vercel.com"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * time.Hour
	router.Use(cors.New(corsConfig))

	// Authentication
	router.GET("/api/ping", handler.PingHandler)
	router.POST("/api/auth/login", handler.LoginHandler)
	router.POST("/api/auth/register", handler.RegisterHandler)
	router.GET("/api/auth/refresh",middleware.AuthMiddleware(), handler.GetAccessToken)
	router.POST("/api/auth/logout",middleware.AuthMiddleware(), handler.LogoutHandler)

	// Paste
	router.POST("/api/paste", middleware.InjectOptionalUserID(), handler.CreatePasteHandler)
	router.GET("/api/paste/:id", handler.GetPasteHandler)
	router.GET("/api/paste/mine", middleware.AuthMiddleware(), handler.GetMyPastesHandler)
	router.DELETE("/api/paste/:id", middleware.AuthMiddleware(), handler.DeletePasteHandler)

	// User
	router.DELETE("/api/users/me", middleware.AuthMiddleware(), handler.DeleteUserHandler)
	router.GET("/api/users/me", middleware.AuthMiddleware(), handler.GetMeHandler)
	return router
}
