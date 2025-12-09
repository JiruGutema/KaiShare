package handler

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jirugutema/kaishare/internal/repository"
	"github.com/jirugutema/kaishare/internal/service"
	"github.com/jirugutema/kaishare/pkg"
)

func DeleteUserHandler(ctx *gin.Context) {
	userID, err := pkg.GetUUIDFromGinContextParam(ctx, "userID")
	if err != nil {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized Access!",
		})
		return
	}

	user, err := service.DeleteUser(userID)
	fmt.Println(err)
	if errors.Is(err, service.ErrUserNotExist) {

		ctx.JSON(400, gin.H{
			"error": "User doesn't exist!",
		})
		return
	}
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "error while deleting the user",
		})
		return
	}
	pkg.ClearCookies(ctx, domainHost)

	ctx.JSON(200, gin.H{
		"success": true,
		"user":    user,
	})
}

func GetMeHandler(ctx *gin.Context) {
	userID, err := pkg.GetUUIDFromGinContextParam(ctx, "userID")
	if err != nil {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized Access!",
		})
		return
	}

	user, _, err := repository.GetUserByID(userID)
	if err != nil {
		ctx.JSON(401, gin.H{
			"error": "Error retrieving user information!",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"user":    user,
	})
}
