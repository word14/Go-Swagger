package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitSwagger(app *gin.Engine) {
	swaggerConfig := &ginSwagger.Config{
		URL: "http://localhost:3000/swagger/doc.json", //The url pointing to API definition
	}
	app.GET("/swagger/*any", ginSwagger.CustomWrapHandler(swaggerConfig, swaggerFiles.Handler))
}
