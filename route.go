package main

import (
	"fmt"

	"./controller"
	"./middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	app.GET("/", func(context *gin.Context) {
		fmt.Println("masok!!")
	}, controller.HomeGet, func(context *gin.Context) {
		fmt.Println("masok 2 !!")
	})

	app.GET("/test",
		// middleware.JwtMiddleware,
		controller.ProductSuperCategoryList)
	app.GET("/test/:slug", middleware.JwtMiddleware, controller.ProductSuperCategoryGet)
	app.POST("/", controller.HomePost)
	app.POST("/login", controller.Login)
}
