package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jirugutema/gopastebin/internal/dto"
	"github.com/jirugutema/gopastebin/internal/service"
)

func LoginHandler(ctx *gin.Context) {
	var user dto.LoginDTO
	err := ctx.BindJSON(&user)
	fmt.Println(user)
	if err != nil {
		ctx.JSON(400, "Error getting user information")
		return
	}
	loginResponse, accessToken, refreshToken, err := service.LoginService(user)

	if err == sql.ErrNoRows {
		ctx.JSON(404, gin.H{"error": "User does not exist"})
		return
	}
	if errors.Is(err, service.ErrEmailOrPassword) {

		ctx.JSON(401, gin.H{"error": "wrong email or password"})
		return
	}
	if err != nil {
		ctx.JSON(401, "Login failed")
		return
	}

	ctx.SetCookie("access_token", accessToken, 3600, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, int((time.Hour * 24 * 7).Seconds()), "/", "localhost", false, true)
	ctx.JSON(200, gin.H{
		"user": loginResponse,
	})
}

func RegisterHandler(ctx *gin.Context) {
	var user dto.RegisterDTO

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, "Error getting user information")
		return
	}
	resp, accessToken, refreshToken, err := service.RegisterServices(user)

	if errors.Is(err, service.ErrUserExists) {
		ctx.JSON(409, gin.H{
			"Error": "User already exists",
		})
		return
	}

	if errors.Is(err, service.ErrIDGenerating) {
		ctx.JSON(409, gin.H{
			"Error": "Internal Error! Please, try again!",
		})
		return
	}
	if err != nil {
		ctx.JSON(401, gin.H{"error": "Registration failed"})
		return
	}

	ctx.SetCookie("access_token", accessToken, 3600, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, int((time.Hour * 24 * 7).Seconds()), "/", "localhost", false, true)

	ctx.JSON(200, gin.H{
		"user": resp,
	})
}

func GetAccessToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(400, "Error getting refresh token")
		return
	}

	token, err := service.GetAccessTokenService(refreshToken)
	if err != nil {
		ctx.JSON(400, "Error getting refresh token")
		return
	}

	ctx.SetCookie("access_token", token, 3600, "/", "localhost", false, true)
	ctx.JSON(200, gin.H{
		"success": true,
	})
}

func LogoutHandler(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.JSON(200, gin.H{
		"success": true,
	})
}
