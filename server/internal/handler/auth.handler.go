package handler

import (
	"database/sql"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jirugutema/kaishare/internal/config"
	"github.com/jirugutema/kaishare/internal/dto"
	"github.com/jirugutema/kaishare/internal/service"
	"github.com/jirugutema/kaishare/pkg"
)

var domainHost = config.LoadConfig().Domain

func LoginHandler(ctx *gin.Context) {
	var user dto.LoginDTO
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid body"})
		return
	}

	loginResp, accessToken, refreshToken, err := service.LoginService(user)

	if err == sql.ErrNoRows {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}
	if errors.Is(err, service.ErrEmailOrPassword) {
		ctx.JSON(401, gin.H{"error": "Wrong email or password"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Login failed"})
		return
	}

	pkg.SetAuthCookie(ctx, "access_token", accessToken, domainHost, 3600, true)
	pkg.SetAuthCookie(ctx, "refresh_token", refreshToken, domainHost, int((time.Hour*24*7).Seconds()), true)
	pkg.SetAuthCookie(ctx, "logged_in", "1", domainHost, 3600, false)

	ctx.JSON(200, gin.H{"user": loginResp})
}

func RegisterHandler(ctx *gin.Context) {
	var user dto.RegisterDTO
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid body"})
		return
	}

	res, accessToken, refreshToken, err := service.RegisterService(user)

	if errors.Is(err, service.ErrUserExists) {
		ctx.JSON(409, gin.H{"error": "User already exists"})
		return
	}
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Registration failed"})
		return
	}

	pkg.SetAuthCookie(ctx, "access_token", accessToken, domainHost, 3600, true)
	pkg.SetAuthCookie(ctx, "refresh_token", refreshToken, domainHost, int((time.Hour*24*7).Seconds()), true)
	pkg.SetAuthCookie(ctx, "logged_in", "1", domainHost, 3600, false)

	ctx.JSON(200, gin.H{"user": res})
}

func GetAccessToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		pkg.ClearCookies(ctx, domainHost)
		ctx.JSON(401, gin.H{"error": "Refresh token missing"})
		return
	}

	token, err := service.GetAccessTokenService(refreshToken)
	if errors.Is(err, service.ErrValidatingToken) {
		pkg.ClearCookies(ctx, domainHost)
		ctx.JSON(401, gin.H{"error": "Invalid refresh token"})
		return
	}
	if err != nil {
		pkg.ClearCookies(ctx, domainHost)
		ctx.JSON(500, gin.H{"error": "Token refresh failed"})
		return
	}

	pkg.SetAuthCookie(ctx, "access_token", token, domainHost, 3600, true)
	pkg.SetAuthCookie(ctx, "logged_in", "1", domainHost, 3600, false)

	ctx.JSON(200, gin.H{"success": true})
}

func LogoutHandler(ctx *gin.Context) {
	pkg.ClearCookies(ctx, domainHost)
	ctx.JSON(200, gin.H{"success": true})
}
