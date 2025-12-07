package pkg

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

var (
	IsProd = os.Getenv("GO_ENV") == "production"
)

func SetAuthCookie(ctx *gin.Context, name, value, domain string, maxAge int, httpOnly bool) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Domain:   domain,
		MaxAge:   maxAge,
		HttpOnly: httpOnly,
	}

	if IsProd {
		cookie.Secure = true
		cookie.SameSite = http.SameSiteNoneMode
	} else {
		cookie.Secure = false
		cookie.SameSite = http.SameSiteLaxMode
	}

	http.SetCookie(ctx.Writer, cookie)
}

func ClearCookie(ctx *gin.Context, name, domain string) {
	c := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		Domain:   domain,
		MaxAge:   -1,
		HttpOnly: true,
	}

	if IsProd {
		c.Secure = true
		c.SameSite = http.SameSiteNoneMode
	}

	http.SetCookie(ctx.Writer, c)
}

func ClearCookies(ctx *gin.Context, domain string) {
	ClearCookie(ctx, "access_token", domain)
	ClearCookie(ctx, "refresh_token", domain)
	ClearCookie(ctx, "logged_in", domain)
}
