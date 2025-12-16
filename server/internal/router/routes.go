// Package routes contains all route definitions for the API server
package routes

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jirugutema/kaishare/internal/handler"
	"github.com/jirugutema/kaishare/internal/middleware"
)

var IsProd = os.Getenv("GO_ENV") == "production"

func Routes() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	store := cookie.NewStore([]byte("secret")) 
	if IsProd {
		store.Options(sessions.Options{
			Path:     "localhost",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
		})
	} else {
		store.Options(sessions.Options{
			Path:     "localhost",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   false,               
			SameSite: http.SameSiteLaxMode,
		})
	}
	router.Use(sessions.Sessions("unlockedpastes", store))
	// router.Use(pkg.Logger)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "https://kai-share.vercel.app"}
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
	router.GET("/api/auth/refresh", middleware.GetAccessTokenMiddleware(), handler.GetAccessToken)
	router.POST("/api/auth/logout", middleware.AuthMiddleware(), handler.LogoutHandler)
	router.GET("/api/auth/check", middleware.AuthMiddleware(), handler.PingMe)
	router.GET("/api/auth/inject", middleware.InjectUserInformationFromParsedToken(), handler.GetMeHandler)

	// Paste
	router.POST("/api/paste", middleware.InjectOptionalUserID(), handler.CreatePasteHandler)
	router.GET("/api/paste/:id", middleware.InjectOptionalUserID(), handler.GetPasteHandler)
	router.GET("/api/paste/mine", middleware.AuthMiddleware(), handler.GetMyPastesHandler)
	router.DELETE("/api/paste/:id", middleware.AuthMiddleware(), handler.DeletePasteHandler)

	// User
	router.DELETE("/api/users/me", middleware.AuthMiddleware(), handler.DeleteUserHandler)
	router.GET("/api/users/me", middleware.AuthMiddleware(), handler.GetMeHandler)
	return router
}
