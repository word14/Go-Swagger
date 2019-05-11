package controller

import (
	"errors"
	"fmt"

	"../model"
	"github.com/gin-gonic/gin"
)

func HomeGet(context *gin.Context) {
	// for i := 10; i > -1; i-- {
	// 	fmt.Println(7 / i)
	// 	fmt.Println("Kelvin")
	// }
	// employee := model.Employee{
	// 	Name:   "Yusuf",
	// 	Salary: 100,
	// }
	fmt.Println("kelvin")
	// context.JSON(200, employee)
	context.Error(errors.New("error guys!!"))
}

func HomePost(context *gin.Context) {
	var body model.Employee
	context.Bind(&body)

	// contoh := Test{nama: body.nama}
	// db.Create(&contoh)

	context.JSON(200, body)
}
