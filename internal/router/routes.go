// Package routes contains all route definitions for the API server
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jirugutema/gopastebin/internal/handler"
	"github.com/jirugutema/gopastebin/internal/middleware"
	"github.com/jirugutema/gopastebin/pkg"
)

func Routes() *gin.Engine {
	router := gin.Default()
	router.Use(pkg.Logger)

	// Authentication
	router.GET("/ping", middleware.AuthMiddleware(), handler.PingHandler)
	router.POST("/auth/login", handler.LoginHandler)
	router.POST("/auth/register", handler.RegisterHandler)
	router.GET("/auth/refresh", handler.GetAccessToken)
	router.GET("/auth/logout", handler.LogoutHandler)

	// Paste
	router.POST("/paste", middleware.InjectOptionalUserID(), handler.CreatePasteHandler)
	router.GET("/paste/:id" ,handler.GetPasteHandler)
	return router
}

