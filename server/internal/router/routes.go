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

func Routes() *gin.Engine {
	isProd := os.Getenv("GO_ENV") == "production"

	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		sessionSecret = "dev-secret"
	}

	store := cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   isProd,
		SameSite: func() http.SameSite {
			if isProd {
				return http.SameSiteNoneMode
			}
			return http.SameSiteLaxMode
		}(),
	})

	router.Use(sessions.Sessions("unlockedpastes", store))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		"http://localhost:3000",
		"https://kai-share.vercel.app",
	}
	corsConfig.AllowMethods = []string{
		"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
	}
	corsConfig.AllowHeaders = []string{
		"Origin",
		"Content-Type",
		"Authorization",
		"Accept",
		"User-Agent",
		"Cache-Control",
		"Pragma",
	}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * time.Hour

	router.Use(cors.New(corsConfig))

	router.Use(middleware.RateLimiter(200))

	api := router.Group("/api")

	api.GET("/ping", handler.PingHandler)

	auth := api.Group("/auth")
	auth.Use(middleware.RateLimiter(10))
	{
		auth.POST("/login", handler.LoginHandler)
		auth.POST("/register", handler.RegisterHandler)
		auth.GET("/refresh", middleware.GetAccessTokenMiddleware(), handler.GetAccessToken)
		auth.POST("/logout", middleware.AuthMiddleware(), handler.LogoutHandler)
		auth.GET("/check", middleware.AuthMiddleware(), handler.PingMe)
		auth.GET("/inject", middleware.InjectUserInformationFromParsedToken(), handler.GetMeHandler)
	}

	paste := api.Group("/paste")
	{
		paste.POST("", middleware.InjectOptionalUserID(), handler.CreatePasteHandler)
		paste.GET("/:id", middleware.InjectOptionalUserID(), handler.GetPasteHandler)
		paste.GET("/mine", middleware.AuthMiddleware(), handler.GetMyPastesHandler)
		paste.DELETE("/:id", middleware.AuthMiddleware(), handler.DeletePasteHandler)
		paste.PUT("/:id", middleware.AuthMiddleware(), handler.UpdatePasteHandler)
	}

	user := api.Group("/users")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/me", handler.GetMeHandler)
		user.DELETE("/me", handler.DeleteUserHandler)
	}

	return router
}
