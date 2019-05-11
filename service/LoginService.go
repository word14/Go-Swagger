package service

import (
	"os"
	"time"

	"../helper"
	"../model"
	"../repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(body *model.LoginRequest, context *gin.Context) {
	customer := repository.FindByEmail(&body.Email)
	if helper.IsSame(body.Password, customer.Password) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uuid": customer.Uuid,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})
		secret := []byte(os.Getenv("TOKEN_ECRYPT"))
		tokenString, _ := token.SignedString(secret)
		context.JSON(200, gin.H{
			"token": tokenString,
		})
	} else {
		context.JSON(422, gin.H{
			"message": "Invalid username or password",
		})
	}
}
