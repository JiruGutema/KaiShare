package handler

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jirugutema/kaishare/internal/dto"
	"github.com/jirugutema/kaishare/internal/service"
	"github.com/jirugutema/kaishare/pkg"
)

func CreatePasteHandler(ctx *gin.Context) {
	var paste dto.PasteDTO

	userID, err := pkg.GetUUIDFromGinContextParam(ctx, "userID")
	if err == nil {
		paste.UserID = &userID
	}
	err = ctx.BindJSON(&paste)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Error creating paste"})
		return
	}

	pasteID, err := service.CreatePasteService(paste)
	if errors.Is(err, service.ErrUserNotExist) {
		ctx.JSON(404, gin.H{"error": "user associated with the paste doesn't exist"})
		return
	}
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Error creating paste"})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"pasteId": pasteID,
	})
}

func GetPasteHandler(ctx *gin.Context) {
	pasteID, err := uuid.Parse(ctx.Param("id"))
	pastePassword := ctx.Query("password")

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Error getting paste id"})
		return
	}
	paste, err := service.GetPasteService(pasteID)

	if errors.Is(err, service.ErrPasteExpired) {

		ctx.JSON(400, gin.H{"error": "This paste has been expired"})
		return
	}
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Error retrieving paste"})
		fmt.Println(err)
		return
	}

	if paste.Password != nil {

		match := pkg.ComparePasswordHash(pastePassword, *paste.Password)

		if !match {
			emptyPasteResponse := dto.PasteResponse{}
			emptyPasteResponse.RequiresPassword = true
			ctx.JSON(401, gin.H{"error": "wrong password is provided for the paste", "paste": emptyPasteResponse})
			fmt.Println(err)
			return
		}
	}

	ctx.JSON(200, gin.H{
		"paste": paste,
	})
}

func GetMyPastesHandler(ctx *gin.Context) {
	userID, err := pkg.GetUUIDFromGinContextParam(ctx, "userID")
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "error retrieving user pastes",
		})
		return
	}

	myPastes, err := service.GetMyPastesService(userID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "error retrieving user pastes",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"succes": true,
		"pastes": myPastes,
	})
}

func DeletePasteHandler(ctx *gin.Context) {
	pasteID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "unable to get pasteID",
		})

		return
	}
	userID, err := pkg.GetUUIDFromGinContextParam(ctx, "userID")
	if err != nil {
		ctx.JSON(401, gin.H{
			"error": "unable to delete paste. invalid user information!",
		})
		return
	}
	e := service.DeletePasteService(pasteID, userID)

	if errors.Is(e, service.ErrCantDeletePaste) {

		ctx.JSON(401, gin.H{
			"error": "unauthorized access",
		})
		return
	}
	if e != nil {
		ctx.JSON(400, gin.H{
			"error": "unable to get delete paste",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
	})
}
