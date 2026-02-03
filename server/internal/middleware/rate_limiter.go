package middleware

import (
	"go.uber.org/ratelimit"

	"github.com/gin-gonic/gin"
)

func RateLimiter(rps int) gin.HandlerFunc {
	limiter := ratelimit.New(rps)

	return func(c *gin.Context) {
		limiter.Take()
		c.Next()
	}
}
