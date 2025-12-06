package pkg

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUUIDFromGinContextParam(c *gin.Context, key string) (uuid.UUID, error) {
	userIDValue, exists := c.Get(key)
	if !exists {
		return uuid.Nil, fmt.Errorf("userID not found in context")
	}

	userIDStr, ok := userIDValue.(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid userID type in context")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID format")
	}

	return userID, nil
}
