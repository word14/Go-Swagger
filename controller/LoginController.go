package controller

import (
	"../model"
	"../service"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var body model.LoginRequest
	context.Bind(&body)
	service.Login(&body, context)
}

func Register(context *gin.Context) {
	var body model.RegisterRequest
	context.Bind(&body)
	service.Register(&body, context)
}
