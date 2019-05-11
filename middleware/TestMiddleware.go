package middleware

import (
	"github.com/gin-gonic/gin"
)

func TestMiddleware(nama string) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		println("TestMiddleware " + nama)
	}
}

func handleError(context *gin.Context) {
	if r := recover(); r != nil {
		context.JSON(500, gin.H{
			"message": r,
			"success": false,
		})
	} else if len(context.Errors) > 0 {
		context.JSON(422, gin.H{
			"message": context.Errors[0].Error(),
			"success": false,
		})
	}
}

func ErrorMiddleware(context *gin.Context) {
	defer handleError(context)
	context.Next()
}
