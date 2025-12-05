package pkg

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger(context *gin.Context) {
	fmt.Println("Request Body:", context.Request.Body)
	fmt.Println("Request Method:", context.Request.Method)
	fmt.Println("Request URL:", context.Request.URL)
	fmt.Println("Request Host:", context.Request.Host)
	fmt.Println("Request Remote Address:", context.Request.RemoteAddr)
	fmt.Println("Request Content Length:", context.Request.ContentLength)
	fmt.Println("Request Transfer Encoding:", context.Request.TransferEncoding)
	fmt.Println("Request Proto:", context.Request.Proto)
}
