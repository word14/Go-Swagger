package service

import (
	"../model"
	"github.com/gin-gonic/gin"
)

func Register(body *model.RegisterRequest, context *gin.Context) {
	// password, _ := helper.Hash(&body.Password)
	// customer := model.Customer{
	// 	Email:    body.Email,
	// 	Password: password,
	// 	Name:     body.Name}
	// result := repository.
}
