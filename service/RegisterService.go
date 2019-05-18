package service

import (
	"../helper"
	"../model"
	"../repository"
	"github.com/gin-gonic/gin"
)

func Register(body *model.RegisterRequest, context *gin.Context) {
	password, _ := helper.Hash(&body.Password)
	customer := model.Customer{
		Uuid:     "nil",
		Email:    body.Email,
		Password: password,
		Name:     body.Name}
	repository.Create(&customer)
}
