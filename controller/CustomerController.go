package controller

import (
	"errors"

	"../service"
	"github.com/gin-gonic/gin"
)

func FindByEmail(context *gin.Context) {
	email := context.Param("email")
	println("+=============================+")
	println(email)
	println("-=============================-")
	result := service.FindByEmail(&email)
	if result != nil {
		context.JSON(200, result)
	} else {
		context.Error(errors.New("Not Found"))
	}
}
