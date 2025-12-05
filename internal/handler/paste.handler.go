package handler

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jirugutema/gopastebin/internal/dto"
	"github.com/jirugutema/gopastebin/internal/service"
	"github.com/jirugutema/gopastebin/pkg"
)

func CreatePasteHandler(ctx *gin.Context) {
	var paste dto.PasteDTO

	userID, err := pkg.GetUserIDFromContext(ctx)
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
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Error getting paste id"})
		return
	}
	paste, err := service.GetPasteService(pasteID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Error retrieving paste"})
		fmt.Println(err)
		return
	}

	ctx.JSON(200, gin.H{
		"paste": paste,
	})
}
