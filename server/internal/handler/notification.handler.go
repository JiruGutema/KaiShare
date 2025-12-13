package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jirugutema/kaishare/internal/dto"
	"github.com/jirugutema/kaishare/internal/service"
)

func NotificationHandler(ctx *gin.Context) {
	var notification dto.CreateNotificationDTO

	err := ctx.BindJSON(&notification)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "error getting notification information",
		})
	}

	notificationID, err := service.CreateNotificationService(notification)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "error getting notification information",
		})
	}

	ctx.JSON(200, gin.H{
		"success":        true,
		"notificationId": notificationID,
	})
}
