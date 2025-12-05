// Package middleware is used to intercept the request and check and change the rquest
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jirugutema/gopastebin/pkg"
)

func InjectOptionalUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("access_token")
		if err != nil {
			c.Next()
			return
		}

		token, err := pkg.ValidateJWT(tokenStr)
		if err != nil || !token.Valid {
			c.Next()
			return
		}

		claims := pkg.GetClaims(token)
		c.Set("userID", claims["userID"])

		c.Next()
	}
}
