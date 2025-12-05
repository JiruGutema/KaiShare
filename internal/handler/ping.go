// Package handler contains all route handlers for the API server
package handler

import (
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Pong"})
}
