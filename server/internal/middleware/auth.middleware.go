// Package middleware is used to intercept the request and check and change the rquest
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jirugutema/kaishare/pkg"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("access_token")
		if err != nil {

			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return

		}

		token, err := pkg.ValidateJWT(tokenStr)
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		claims := pkg.GetClaims(token)
		c.Set("userID", claims["userID"])


		c.Next()
	}
}
